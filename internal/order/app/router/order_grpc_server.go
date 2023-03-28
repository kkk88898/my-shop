package router

import (
	"context"
	order "myshop/internal/order/usecases/order"

	gen "myshop/proto/gen/order"

	"github.com/google/wire"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ gen.OrderServiceServer = (*orderGRPCServer)(nil)

var OrderGRPCServerSet = wire.NewSet(NewOrderGRPCServer)

type orderGRPCServer struct {
	gen.UnimplementedOrderServiceServer
	uc order.UseCase
}

func NewOrderGRPCServer(
	grpcServer *grpc.Server,
	uc order.UseCase,
) gen.OrderServiceServer {
	svc := orderGRPCServer{
		uc: uc,
	}

	gen.RegisterOrderServiceServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *orderGRPCServer) GetOrderById(
	ctx context.Context,
	request *gen.GetOrderByIdRequest,
) (*gen.GetOrderByIdResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetUser")

	if request.OrderId == "1" {
		return &gen.GetOrderByIdResponse{
			Order: &gen.OrderDto{
				OrderId:    "1",
				Price:      10,
				Type:       1,
				Image:      "http://ftp.simbalink.cn/123.png",
				CreateTime: "2023-03-11 12:21:21",
			}}, nil
	} else {
		return &gen.GetOrderByIdResponse{}, nil
	}
}

func (g *orderGRPCServer) DelOrderById(
	ctx context.Context,
	request *gen.DelOrderRequest,
) (*gen.DelOrderResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetUser")

	if request.OrderId == "1" {
		return &gen.DelOrderResponse{
			Count: 1,
		}, nil
	} else {
		return &gen.DelOrderResponse{
			Count: 0,
		}, nil
	}
}
