package domain

import (
	"context"
)

type (
	OrderRepo interface {
		DelOrderById(c context.Context, orerId string) (*OrderDto, error)
		GetOrderById(c context.Context, orerId string) (*OrderDto, error)
	}
)
