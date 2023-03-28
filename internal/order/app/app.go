package app

import (
	"myshop/cmd/order/config"

	orderUC "myshop/internal/order/usecases/order"

	gen "myshop/proto/gen/order"
)

type App struct {
	Cfg             *config.Config
	UC              orderUC.UseCase
	OrderGRPCServer gen.OrderServiceServer
}

func New(
	cfg *config.Config,
	uc orderUC.UseCase,
	orderGRPCServer gen.OrderServiceServer,
) *App {
	return &App{
		Cfg:             cfg,
		UC:              uc,
		OrderGRPCServer: orderGRPCServer,
	}
}
