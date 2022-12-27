package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func Connect(ctx context.Context) *Redis {
	r := &Redis{
		ctx: ctx,
	}

	log.Println("Connecting to REDIS...")
	if err := r.connectToRedis(); err != nil {
		log.Println("Error connecting to REDIS: ", err)
	}
	log.Println("Connected to REDIS")

	return r
}

func (r *Redis) CheckConnection() error {
	if r.client == nil {
		return errors.New("")
	}
	_, err := r.client.Ping(r.ctx).Result()

	return err
}

func (r *Redis) Close() {
	log.Println("Disconnect from REDIS...")
	if r.client != nil {
		r.client.Close()
	}

	log.Println("Disconnected from REDIS")
}

func (r *Redis) connectToRedis() error {
	var client *redis.Client
	var err error

	for i := 1; i <= 3; i++ {
		client, err = r.connect()
		if err == nil {
			break
		}

		log.Println("Error connect: ", err)
		log.Printf("Next try after %d seconds...", i*5)
		time.Sleep(time.Duration(i*5) * time.Second)
	}

	if err != nil {
		return err
	}
	if client == nil {
		return errors.New("error creating client")
	}

	r.client = client

	return nil
}

func (r *Redis) connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping(r.ctx).Result(); err != nil {
		return nil, err
	}

	return client, nil
}
