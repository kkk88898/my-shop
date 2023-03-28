package app

import (
	"myshop/cmd/user/config"

	userUC "myshop/internal/user/usecases/users"

	gen "myshop/proto/gen/user"
)

type App struct {
	Cfg            *config.Config
	UC             userUC.UseCase
	UserGRPCServer gen.UserServer
}

func New(
	cfg *config.Config,
	uc userUC.UseCase,
	userGRPCServer gen.UserServer,
) *App {
	return &App{
		Cfg:            cfg,
		UC:             uc,
		UserGRPCServer: userGRPCServer,
	}
}
