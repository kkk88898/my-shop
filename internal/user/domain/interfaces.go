package domain

import (
	"context"
)

type (
	UserRepo interface {
		GetAll(context.Context) ([]*ItemTypeDto, error)
		GetByTypes(context.Context, []string) ([]*ItemDto, error)
	}
)
