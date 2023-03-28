//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"myshop/cmd/order/config"
	"myshop/internal/order/app/router"
	"myshop/internal/order/infras/repo"
	orderUC "myshop/internal/order/usecases/order"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
) (*App, error) {
	panic(wire.Build(
		New,
		router.OrderGRPCServerSet,
		repo.RepositorySet,
		orderUC.UseCaseSet,
	))
}
