package main

import (
	"context"
	"log"
	"time"
	"user_service/redis"
	"user_service/server"
	"user_service/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rdb := redis.Connect(ctx)
	defer rdb.Close()

	fillRDB(rdb)

	svc := service.New(rdb)
	go server.MakeAndRunGRPCAuthServer(":5000", svc)
	conn := service.MakeGRPCAuthConn()
	defer conn.Close()

	{ // making a request to the auth_service
		time.Sleep(time.Second * 5)
		service.AuthUser(ctx, conn, "example@example.com", "PasswordHash*&^&*&^")
	} // conn is there but without time.Sleep there is an error: could not auth: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 172.18.0.3:5000: connect: connection refused"

	<-ctx.Done()
}

func fillRDB(rdb *redis.Redis) {
	if _, err := rdb.CreateUser(&redis.User{Email: "test@example.com", PasswordHash: "test^&^*%"}); err != nil {
		log.Println("fillRDB:", err)
	}
}
