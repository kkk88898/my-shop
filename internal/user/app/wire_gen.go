// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"google.golang.org/grpc"
	"myshop/cmd/user/config"
	"myshop/internal/user/app/router"
	"myshop/internal/user/infras/repo"
	"myshop/internal/user/usecases/users"
)

// Injectors from wire.go:

func InitApp(cfg *config.Config, grpcServer *grpc.Server) (*App, error) {
	userRepo := repo.NewOrderRepo()
	useCase := users.NewService(userRepo)
	userServer := router.NewUserGRPCServer(grpcServer, useCase)
	app := New(cfg, useCase, userServer)
	return app, nil
}
