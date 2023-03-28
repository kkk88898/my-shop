package users

import (
	"context"
	"strings"

	"myshop/internal/user/domain"

	"github.com/google/wire"
	"github.com/pkg/errors"
)

var _ UseCase = (*service)(nil)

var UseCaseSet = wire.NewSet(NewService)

type service struct {
	repo domain.UserRepo
}

func NewService(repo domain.UserRepo) UseCase {
	return &service{
		repo: repo,
	}
}

func (s *service) GetItemTypes(ctx context.Context) ([]*domain.ItemTypeDto, error) {
	results, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemTypes")
	}

	return results, nil
}

func (s *service) GetItemsByType(ctx context.Context, itemTypes string) ([]*domain.ItemDto, error) {
	types := strings.Split(itemTypes, ",")

	results, err := s.repo.GetByTypes(ctx, types)
	if err != nil {
		return nil, errors.Wrap(err, "service.GetItemsByType")
	}

	return results, nil
}
