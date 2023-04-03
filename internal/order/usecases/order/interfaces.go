package order

import (
	"context"

	"myshop/internal/order/domain"
)

type (
	Repo interface {
		UpdateById(c context.Context, orderId string) (*domain.OrderDto, error)
		SelectById(c context.Context, orderId string) (*domain.OrderDto, error)
	}

	UseCase interface {
		DelOrderById(context.Context, string) (*domain.OrderDto, error)
		GetOrderById(context.Context, string) (*domain.OrderDto, error)
	}
)
