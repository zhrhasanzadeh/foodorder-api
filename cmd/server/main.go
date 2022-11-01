package main

import (
	httpHandler "api-orderfood/internal/handler"
	p "api-orderfood/internal/repository/postgres"
	r "api-orderfood/internal/repository/redis"
	service "api-orderfood/internal/service"
	config "api-orderfood/pkg"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cache := redis.NewClient(&redis.Options{
		Addr:     c.RedisURL,
		Password: "",
		DB:       0,
	})

	redisRepo := r.NewRedisRepo(cache, ctx)

	conn, err := pgx.Connect(context.Background(), c.PostgresURL)
	postgresRepo := p.NewPostgresRepository(conn)

	s := service.NewService(&c, postgresRepo, redisRepo)

	httpHandler := httpHandler.NewHttpHandler(s)
	http.HandleFunc("/order/get", httpHandler.GetOrder)
	http.HandleFunc("/order/add", httpHandler.AddOrder)
	log.Println("starting server at" + c.Port)
	err = http.ListenAndServe(c.Port, nil)
	if err != nil {
		panic(err)
	}
}
