package order

import (
	"context"

	"github.com/google/wire"
	"github.com/pkg/errors"
	"myshop/internal/order/domain"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo Repo
}

func NewService(repo Repo) UseCase {
	return &service{
		repo: repo,
	}
}

func (s *service) DelOrderById(ctx context.Context, orderId string) (*domain.OrderDto, error) {
	orderDto, err := s.repo.UpdateById(ctx, orderId)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}
	return orderDto, nil
}

func (s *service) GetOrderById(ctx context.Context, orderId string) (*domain.OrderDto, error) {
	results, err := s.repo.SelectById(ctx, orderId)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemsByType")
	}
	return results, nil
}
