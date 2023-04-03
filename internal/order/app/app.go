package app

import (
	"myshop/cmd/order/config"
	orderUC "myshop/internal/order/usecases/order"
	"myshop/pkg/mysql8"

	gen "myshop/proto/gen/order"
)

type App struct {
	Cfg             *config.Config
	MYSQL8          mysql8.DBEngine
	UC              orderUC.UseCase
	OrderGRPCServer gen.OrderServiceServer
}

func New(
	cfg *config.Config,
	uc orderUC.UseCase,
	mysql8 mysql8.DBEngine,
	orderGRPCServer gen.OrderServiceServer,
) *App {
	return &App{
		Cfg:             cfg,
		UC:              uc,
		OrderGRPCServer: orderGRPCServer,
		MYSQL8:          mysql8,
	}
}
