package grpc

import (
	svc "auth_service/app/service"
	"auth_service/proto"
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"google.golang.org/grpc"
)

func MakeAndRunGRPCAuthServer(port string, svc svc.Auther) {
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

func MakeGRPCAuthConn() *grpc.ClientConn {
	conn, err := grpc.Dial("user_service:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

type GRPCAuthServer struct {
	svc svc.Auther
	proto.UnimplementedAuthServer
}

func NewGRPCAuthServer(svc svc.Auther) *GRPCAuthServer {
	return &GRPCAuthServer{svc: svc}
}

func (s *GRPCAuthServer) Authorize(ctx context.Context, u *proto.User) (*proto.AuthResponse, error) {
	res := s.svc.Authorize(ctx, u.Email, u.PasswordHash)
	fmt.Printf("Got a user (%+v)\n", u)
	return &proto.AuthResponse{Success: res}, nil
}
