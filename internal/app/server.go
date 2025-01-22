package server

import (
	"google.golang.org/grpc"
	"market/internal/app/getuser"
	"market/internal/app/login"
	"market/internal/app/logout"
	"market/internal/app/updateuser"
)

type Server struct {
	LoginHandler      *login.LoginHandler
	LogoutHandler     *logout.LogoutHandler
	GetUserHandler    *getuser.GetUserHandler
	UpdateUserHandler *updateuser.UpdateUserHandler
}

func NewServer(getUserHandler *getuser.GetUserHandler, loginHandler *login.LoginHandler, logoutHandler *logout.LogoutHandler, updateUserHandler *updateuser.UpdateUserHandler) *Server {
	return &Server{
		GetUserHandler:    getUserHandler,
		LoginHandler:      loginHandler,
		LogoutHandler:     logoutHandler,
		UpdateUserHandler: updateUserHandler,
	}
}

func (s *Server) Register(gRPC *grpc.Server) {
	s.LoginHandler.Register(gRPC)
	s.LogoutHandler.Register(gRPC)
	s.GetUserHandler.Register(gRPC)
	s.UpdateUserHandler.Register(gRPC)
}
