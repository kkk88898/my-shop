package order

import (
	"context"

	"myshop/internal/order/domain"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo domain.OrderRepo
}

func NewService(repo domain.OrderRepo) UseCase {
	return &service{
		repo: repo,
	}
}

func (s *service) DelOrderById(ctx context.Context, orderId string) (*domain.OrderDto, error) {
	results, err := s.repo.DelOrderById(ctx, orderId)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}

	return results, nil
}

func (s *service) GetOrderById(ctx context.Context, orderId string) (*domain.OrderDto, error) {

	results, err := s.repo.GetOrderById(ctx, orderId)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemsByType")
	}

	return results, nil
}
