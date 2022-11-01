package service

import (
	"api-orderfood/internal/model"
	pkg "api-orderfood/pkg"
	"api-orderfood/pkg/utils"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"strconv"
)

type service struct {
	config *pkg.Config
	dbRepo model.Database
	cache  model.Cache
}

func NewService(config *pkg.Config, dbRepo model.Database, cache model.Cache) model.Service {
	return &service{
		config: config,
		dbRepo: dbRepo,
		cache:  cache,
	}
}

func (s service) AddOrder(order model.FoodOrder) error {
	err := s.dbRepo.InsertOrder(order)
	if err != nil {
		return err
	}
	return nil
}

func (s service) GetUser(id int) (order model.FoodOrder, err error) {
	stringId := strconv.Itoa(id)
	f, err := s.cache.Get(stringId)
	if errors.Is(err, redis.Nil) {
		order, err := s.dbRepo.GetOrder(id)
		if err != nil {
			panic(err)
		}
		fmt.Println("found data in db.")
		json, err := utils.StructToJson(order)
		if err != nil {
			panic(err)
		}
		err = s.cache.Set(stringId, json)
		if err != nil {
			panic(err)
		}
		return order, nil
	}
	orderFood, err := utils.JsonToStruct(f)
	fmt.Println("found data in cache.")
	return orderFood, nil
}
