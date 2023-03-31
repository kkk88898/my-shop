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
	"myshop/pkg/mysql8"
	"myshop/pkg/postgres"
)

func InitApp(
	cfg *config.Config,
	dbPGConnStr postgres.DBConnString,
	dbMysql8ConnStr mysql8.DBConnString,
	grpcServer *grpc.Server,
) (*App, func(), error) {
	panic(wire.Build(
		New,
		dbPGEngineFunc,
		dbMysql8EngineFunc,
		router.OrderGRPCServerSet,
		repo.RepositorySet,
		orderUC.UseCaseSet,
	))
}
func dbPGEngineFunc(url postgres.DBConnString) (postgres.DBEngine, func(), error) {
	db, err := postgres.NewPostgresDB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}

func dbMysql8EngineFunc(url mysql8.DBConnString) (mysql8.DBEngine, func(), error) {
	db, err := mysql8.NewMysql8DB(url)
	if err != nil {
		return nil, nil, err
	}
	return db, func() { db.Close() }, nil
}
