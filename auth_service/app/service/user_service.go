package service

import (
	"auth_service/proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func AuthUser(ctx context.Context, conn *grpc.ClientConn, email, password string) {
	c := proto.NewAuthClient(conn)

	r, err := c.Authorize(ctx, &proto.User{Email: email, PasswordHash: password})
	if err != nil {
		log.Fatalf("could not auth: %v", err)
	}
	log.Printf("Auth result (false):(%v)", r.GetSuccess())
}
