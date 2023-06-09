//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"myshop/cmd/user/config"
	"myshop/internal/user/app/router"
	"myshop/internal/user/infras/repo"
	userUC "myshop/internal/user/usecases/users"
)

func InitApp(
	cfg *config.Config,
	grpcServer *grpc.Server,
) (*App, error) {
	panic(wire.Build(
		New,
		router.UserGRPCServerSet,
		repo.RepositorySet,
		userUC.UseCaseSet,
	))
}
