package app

import (
	"myshop/cmd/order/config"
	"myshop/pkg/mysql8"
	"myshop/pkg/postgres"

	orderUC "myshop/internal/order/usecases/order"

	gen "myshop/proto/gen/order"
)

type App struct {
	Cfg             *config.Config
	PG              postgres.DBEngine
	MYSQL8          mysql8.DBEngine
	UC              orderUC.UseCase
	OrderGRPCServer gen.OrderServiceServer
}

func New(
	cfg *config.Config,
	uc orderUC.UseCase,
	pg postgres.DBEngine,
	mysql8 mysql8.DBEngine,
	orderGRPCServer gen.OrderServiceServer,
) *App {
	return &App{
		Cfg:             cfg,
		UC:              uc,
		OrderGRPCServer: orderGRPCServer,
		PG:              pg,
		MYSQL8:          mysql8,
	}
}
