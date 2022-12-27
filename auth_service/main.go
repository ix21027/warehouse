package main

import (
	"auth_service/app/service"
	"auth_service/internal/grpc"
	"auth_service/internal/messaging"
	"context"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	nats := messaging.Default()
	defer nats.Stop()
	grpcConn := grpc.MakeGRPCAuthConn()
	defer grpcConn.Close()

	svc := service.NewAuth(nats.Conn, grpcConn)
	grpc.MakeAndRunGRPCAuthServer(":5000", svc)

	<-ctx.Done()
}
