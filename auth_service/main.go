package main

import (
	"auth_service/app/service"
	"auth_service/internal/messaging"
	"context"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	nats := messaging.Default()
	defer nats.Stop()

	svc := service.New(nats.Conn)
	svc.Authorize()
	<-ctx.Done()
}
