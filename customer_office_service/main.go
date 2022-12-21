package main

import (
	"context"
	"customer_office_service/app/controller"
	"customer_office_service/app/service"
	"customer_office_service/infrastracture/servers/http_server"
	"customer_office_service/infrastracture/servers/nats_server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	natsServer := nats_server.New()
	go natsServer.Run()
	defer natsServer.Stop()

	mainService := service.NewMain(natsServer)

	httpController := controller.New(mainService)
	httpServer := http_server.NewServer(httpController)
	go httpServer.Run()

	<-ctx.Done()
}
