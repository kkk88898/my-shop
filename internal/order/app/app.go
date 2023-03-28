package app

import (
	"myshop/cmd/user/config"

	productUC "myshop/internal/user/usecases/users"

	gen "myshop/proto/gen/user"
)

type App struct {
	Cfg            *config.Config
	UC             productUC.UseCase
	UserGRPCServer gen.UserServer
}

func New(
	cfg *config.Config,
	uc productUC.UseCase,
	userGRPCServer gen.UserServer,
) *App {
	return &App{
		Cfg:            cfg,
		UC:             uc,
		UserGRPCServer: userGRPCServer,
	}
}
