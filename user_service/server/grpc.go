package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_service/proto"
	"user_service/service"
)

func MakeAndRunGRPCAuthServer(port string, svc service.UserService) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAuthServer(s, NewGRPCAuthServer(svc))

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GRPCAuthServer struct {
	proto.UnimplementedAuthServer
	svc service.UserService
}

func NewGRPCAuthServer(svc service.UserService) *GRPCAuthServer {
	return &GRPCAuthServer{
		svc: svc,
	}
}

func (s *GRPCAuthServer) Authorize(_ context.Context, u *proto.User) (*proto.AuthResponse, error) {
	fmt.Printf("Got a user (%+v)\n", u)
	return &proto.AuthResponse{Success: false}, nil
}

func (s *GRPCAuthServer) UserInfo(_ context.Context, in *proto.UserReq) (*proto.UserRes, error) {
	u, err := s.svc.UserInfo(in.Email)
	if err != nil {
		return nil, err
	}
	return &proto.UserRes{Id: u.Id, PasswordHash: u.PasswordHash}, nil
}
