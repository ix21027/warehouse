package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"user_service/nats"
	"user_service/redis"
	"user_service/scylla"
	"user_service/server"
	"user_service/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	s := scylla.CreateAndConnect()
	defer s.Close()
	fillSDB(s)

	rdb := redis.Connect(ctx)
	defer rdb.Close()
	fillRDB(rdb)

	svc := service.New(rdb, s)
	go server.MakeAndRunGRPCAuthServer(":5000", svc)
	conn := service.MakeGRPCAuthConn()
	defer conn.Close()

	{ // making a request to the auth_service
		time.Sleep(time.Second * 5)
		service.AuthUser(ctx, conn, "example@example.com", "PasswordHash*&^&*&^")
	} // conn is there but without time.Sleep there is an error: could not auth: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing dial tcp 172.18.0.3:5000: connect: connection refused"

	nats := nats.Default(svc)
	defer nats.Stop()

	<-ctx.Done()
}

func fillRDB(rdb *redis.Redis) {
	if _, err := rdb.CreateUser(&redis.User{Email: "test@example.com", PasswordHash: "test^&^*%"}); err != nil {
		log.Println("fillRDB:", err)
	}
}

func fillSDB(s *scylla.Scylla) {
	good := s.CreateGood("testName of Good", "test description of good", "test image of good", 100)
	user := s.CreateUser("custName", "custLogin", "password", "customer")
	s.CreateUser("name", "login", "password", "admin")

	for i := 0; i < 10; i++ {
		s.CreateGood("GoodName_"+strconv.Itoa(i), "description1"+strconv.Itoa(i), "test image", 10011+i)
		s.CreateUser("customer_"+strconv.Itoa(i), "login_"+strconv.Itoa(i), "password", "customer")
	}

	log.Println("s.GetGoodByID(good.ID):", s.GetGoodByID(good.ID))
	log.Println("s.GetUserByID(user.ID):", s.GetUserByID(fmt.Sprintf("%s", user.ID)))

	log.Println(`s.GetGoodByName("name1"):`, s.GetGoodByName("GoodName_4"))
	s.UpdateGoodStatusToDeleted(good.ID)

	s.BanUserByID(fmt.Sprintf("%s", user.ID))
	log.Printf(`s.GetUsersByStatus("ban"): `)
	s.GetUsersByStatus("ban")
}
