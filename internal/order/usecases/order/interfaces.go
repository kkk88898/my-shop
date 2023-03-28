package order

import (
	"context"

	"myshop/internal/order/domain"
)

type UseCase interface {
	DelOrderById(context.Context, string) (*domain.OrderDto, error)
	GetOrderById(context.Context, string) (*domain.OrderDto, error)
}
