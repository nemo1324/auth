package auth

import (
	"context"
	"market/protos/gen/go/auth"

	"google.golang.org/grpc"
)

type serverAPI struct {
	auth.UnimplementedAuthServiceServer
}

func Register(gRPC *grpc.Server) {
	auth.RegisterAuthServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *auth.LoginRequest,
) (*auth.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Logout(
	ctx context.Context,
	req *auth.LogoutRequest,
) (*auth.LogoutResponse, error) {
	panic("implement me")
}

func (s *serverAPI) GetUser(
	ctx context.Context,
	req *auth.GetUserRequest,
) (*auth.GetUserResponse, error) {
	panic("implement me")
}

func (s *serverAPI) UpdateUser(
	ctx context.Context,
	req *auth.UpdateUserRequest, // используем правильный тип
) (*auth.UpdateUserResponse, error) { // возвращаем правильный тип
	panic("implement me")
}
