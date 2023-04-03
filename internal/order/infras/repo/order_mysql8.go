package repo

import (
	"context"
	"database/sql"
	"github.com/pingcap/log"
	"myshop/internal/order/domain"
	"myshop/internal/order/infras/mysql8sql/gen"
	"myshop/internal/order/usecases/order"
	"myshop/pkg/mysql8"

	"github.com/google/wire"
)

var _ order.Repo = (*orderMysql8Repo)(nil)

var RepositorySet = wire.NewSet(NewOrderRepo)

type orderMysql8Repo struct {
	mysqldb mysql8.DBEngine
}

// UpdateById implements order.Repo
func (o *orderMysql8Repo) UpdateById(c context.Context, orderId string) (*domain.OrderDto, error) {
	queries := gen.New(o.mysqldb.GetDB())
	query := gen.Querier(queries)
	id, err := query.DelOrderById(c, gen.DelOrderByIdParams{
		DeleteIs: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		OrderID: orderId,
	})
	if err != nil {
		log.Error("删除报错")
	}
	if id <= 0 {
		log.Error("数据不存在")
	}
	return nil, nil
}

// SelectById implements order.Repo
func (o *orderMysql8Repo) SelectById(c context.Context, orderId string) (*domain.OrderDto, error) {
	queries := gen.New(o.mysqldb.GetDB())
	order, err := queries.GetOrderById(c, orderId)
	//query := gen.Querier(queries)
	//order, err := query.GetOrderById(c, orderId)
	if err != nil {
		log.Error("获取报错")
	}
	dto := domain.OrderDto{
		Name:       "",
		Type:       order.OrderType.Int32,
		Price:      0,
		Image:      "",
		CreateTime: "",
		UpdateTime: "",
	}
	return &dto, nil
}

func NewOrderRepo(mysql8 mysql8.DBEngine) order.Repo {
	return &orderMysql8Repo{
		mysqldb: mysql8,
	}
}
