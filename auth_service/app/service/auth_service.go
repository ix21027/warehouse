package service

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

type Auther interface {
	Authorize(context.Context, string, string) bool
}

type AuthService struct {
	nats *nats.Conn
	grpc *grpc.ClientConn
}

func NewAuth(nc *nats.Conn, grpc *grpc.ClientConn) *AuthService {
	return &AuthService{nats: nc, grpc: grpc}
}

func (s *AuthService) Authorize(ctx context.Context, email, passwordHash string) bool {
	_, err := s.nats.Subscribe("authS.authorize", func(m *nats.Msg) {
		fmt.Println(string(m.Data), "gets from authS.Authorize", email, passwordHash)

		reply := "You have been successfully authorized, " + email

		err := s.nats.Publish("customer_officeS.Authorize", []byte(reply))

		if err != nil {
			panic(err)
		}
	})
	AuthUser(ctx, s.grpc, "not_email", "not_password")

	UserInfo(ctx, s.grpc, "test@example.com")

	if err != nil {
		panic(err)
	}
	return true
}
