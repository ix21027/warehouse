package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"user_service/proto"
)

func MakeGRPCAuthConn() *grpc.ClientConn {
	conn, err := grpc.Dial("auth_service:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func AuthUser(ctx context.Context, conn *grpc.ClientConn, email, password string) {
	c := proto.NewAuthClient(conn)

	r, err := c.Authorize(ctx, &proto.User{Email: email, PasswordHash: password})
	if err != nil {
		log.Fatalf("could not auth: %v", err)
	}
	log.Printf("Auth result (true):(%v)", r.GetSuccess())
}
