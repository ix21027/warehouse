package main

import (
	"context"
	"time"
	"user_service/server"
	"user_service/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go server.MakeAndRunGRPCAuthServer(":5000")
	conn := service.MakeGRPCAuthConn()
	defer conn.Close()

	{ // making a request to the auth_service
		time.Sleep(time.Second * 5)
		service.AuthUser(ctx, conn, "example@example.com", "PasswordHash*&^&*&^")
	} // conn is there but without time.Sleep there is an error: could not auth: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 172.18.0.3:5000: connect: connection refused"

	<-ctx.Done()
}
