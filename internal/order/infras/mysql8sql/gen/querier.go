// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package gen

import (
	"context"
)

type Querier interface {
	DelOrderById(ctx context.Context, arg DelOrderByIdParams) (int64, error)
	GetOrderById(ctx context.Context, orderID string) (XfOrder, error)
}

var _ Querier = (*Queries)(nil)
