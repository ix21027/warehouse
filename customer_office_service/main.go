package main

import (
	"context"
	"customer_office_service/app/controller"
	"customer_office_service/app/service"
	"customer_office_service/infrastracture/server"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	//c := http.NewHTTP()
	//defer c.NatsConn.Close()

	//c.SetRouts().ServeHTTP()
	//********************************//
	ctx, _ := context.WithCancel(context.Background())
	s := service.New()
	httpctlr := controller.NewHTTPController(s)
	conf := server.NewConfig()
	server, _ := server.NewHTTPServer(conf, httpctlr)
	server.Run()
	<-ctx.Done()
	//go server.RunNats(controller)
	//go server.RunHTTP(controller)

	fmt.Println(os.Getenv("HTTP_PORT"))
}
