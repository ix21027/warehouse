package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_service/proto"
)

func MakeAndRunGRPCAuthServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAuthServer(s, NewGRPCAuthServer())

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GRPCAuthServer struct {
	proto.UnimplementedAuthServer
}

func NewGRPCAuthServer() *GRPCAuthServer {
	return &GRPCAuthServer{}
}

func (s *GRPCAuthServer) Authorize(_ context.Context, u *proto.User) (*proto.AuthResponse, error) {
	fmt.Printf("Got a user (%+v)\n", u)
	return &proto.AuthResponse{Success: false}, nil
}
