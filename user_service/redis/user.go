package redis

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(id string) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id string) error
}

type User struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

func (r *Redis) CreateUser(user *User) (*User, error) {
	user.Id = uuid.New().String()
	marshaled, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	r.client.HSet(r.ctx, "users", user.Email, marshaled)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Redis) GetUser(email string) (*User, error) {
	val, err := r.client.HGet(r.ctx, "users", email).Result()

	if err != nil {
		return nil, err
	}
	user := &User{}
	err = json.Unmarshal([]byte(val), user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Redis) GetUsers() ([]*User, error) {
	var users []*User
	val, err := r.client.HGetAll(r.ctx, "users").Result()
	if err != nil {
		return nil, err
	}
	for _, item := range val {
		user := &User{}
		err := json.Unmarshal([]byte(item), user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *Redis) UpdateUser(user *User) (*User, error) {
	marshaled, err := json.Marshal(&user)
	if err != nil {
		return nil, err
	}
	r.client.HSet(r.ctx, "users", user.Email, marshaled)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *Redis) DeleteUser(id string) error {
	numDeleted, err := r.client.HDel(r.ctx, "users", id).Result()
	if numDeleted == 0 {
		return errors.New("user to delete not found")
	}
	if err != nil {
		return err
	}
	return nil
}
