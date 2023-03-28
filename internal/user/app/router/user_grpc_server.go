package router

import (
	"context"

	user "myshop/internal/user/usecases/users"
	gen "myshop/proto/gen/user"

	"github.com/google/wire"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var _ gen.UserServer = (*userGRPCServer)(nil)

var UserGRPCServerSet = wire.NewSet(NewUserGRPCServer)

type userGRPCServer struct {
	gen.UnimplementedUserServer
	uc user.UseCase
}

func NewUserGRPCServer(
	grpcServer *grpc.Server,
	uc user.UseCase,
) gen.UserServer {
	svc := userGRPCServer{
		uc: uc,
	}

	gen.RegisterUserServer(grpcServer, &svc)

	reflection.Register(grpcServer)

	return &svc
}

func (g *userGRPCServer) GetUser(
	ctx context.Context,
	request *gen.IdRequest,
) (*gen.UserResponse, error) {
	slog.Info("gRPC client", "http_method", "GET", "http_name", "GetUser")

	if request.Id == "1" {
		return &gen.UserResponse{
			Id:     "1",
			Name:   "tsw",
			Gender: "男",
		}, nil
	} else {
		return &gen.UserResponse{
			Id:     "1",
			Name:   "不存在",
			Gender: "未知",
		}, nil
	}
}
