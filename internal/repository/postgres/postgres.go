package postgres

import (
	"api-orderfood/internal/model"
	"context"
	"github.com/jackc/pgx/v4"
)

type PostgresRepo struct {
	db *pgx.Conn
}

func NewPostgresRepository(db *pgx.Conn) model.Database {
	return &PostgresRepo{db}
}

func (p *PostgresRepo) InsertOrder(order model.FoodOrder) error {
	rows, err := p.db.Query(context.Background(),
		`insert into "Order"(price,title)values($1,$2)`,
		order.Price, order.Title)
	rows.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) GetOrder(id int) (model.FoodOrder, error) {
	var foodOrder model.FoodOrder
	err := p.db.QueryRow(context.Background(), `select * from "Order" where id=$1`, id).Scan(
		&foodOrder.OrderID, &foodOrder.Price, &foodOrder.Title)
	if err != nil {
		return model.FoodOrder{}, err
	}
	return foodOrder, nil
}
