package repo

import (
	"context"

	"myshop/internal/order/domain"

	"github.com/google/wire"
)

var _ domain.OrderRepo = (*orderInMemRepo)(nil)

var RepositorySet = wire.NewSet(NewOrderRepo)

type orderInMemRepo struct {
	itemTypes map[string]*domain.OrderDto
}

// DelOrderById implements domain.OrderRepo
func (*orderInMemRepo) DelOrderById(c context.Context, orerId string) (*domain.OrderDto, error) {
	panic("unimplemented")
}

// GetOrderById implements domain.OrderRepo
func (*orderInMemRepo) GetOrderById(c context.Context, orerId string) (*domain.OrderDto, error) {
	panic("unimplemented")
}

func NewOrderRepo() domain.OrderRepo {
	return &orderInMemRepo{
		itemTypes: map[string]*domain.OrderDto{
			"CAPPUCCINO": {
				Name:       "CAPPUCCINO",
				Type:       0,
				Price:      4.5,
				Image:      "img/CAPPUCCINO.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
			"COFFEE_BLACK": {
				Name:       "COFFEE_BLACK",
				Type:       1,
				Price:      3,
				Image:      "img/COFFEE_BLACK.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
			"COFFEE_WITH_ROOM": {
				Name:       "COFFEE_WITH_ROOM",
				Type:       2,
				Price:      3,
				Image:      "img/COFFEE_WITH_ROOM.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
			"ESPRESSO": {
				Name:       "ESPRESSO",
				Type:       3,
				Price:      3.5,
				Image:      "img/ESPRESSO.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
			"ESPRESSO_DOUBLE": {
				Name:       "ESPRESSO_DOUBLE",
				Type:       4,
				Price:      4.5,
				Image:      "img/ESPRESSO_DOUBLE.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
			"LATTE": {
				Name:       "LATTE",
				Type:       5,
				Price:      4.5,
				Image:      "img/LATTE.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
			"CAKEPOP": {
				Name:       "CAKEPOP",
				Type:       6,
				Price:      2.5,
				Image:      "img/CAKEPOP.png",
				CreateTime: "2023-03-28 15:21:12",
				UpdateTime: "2023-03-28 15:21:12",
			},
		},
	}
}
