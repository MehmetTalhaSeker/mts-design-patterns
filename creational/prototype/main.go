package main

import (
	"fmt"
	"time"
)

type Model interface {
	Clone() Model
}

type model struct {
	createdAt time.Time
	updatedAt time.Time
}

type UserModel struct {
	model
	Username string
	Email    string
}

func (r *UserModel) Clone() Model {
	return &UserModel{
		Username: r.Username,
		Email:    r.Email,
	}
}

func main() {
	u := &UserModel{
		model: model{
			createdAt: time.Now(),
			updatedAt: time.Now(),
		},
		Username: "username",
		Email:    "user@hemail.com",
	}

	cu := u.Clone()

	a, ok := cu.(*UserModel)
	if !ok {
		return
	}

	a.createdAt = time.Now().Add(time.Minute * 149)
	a.updatedAt = time.Now().Add(time.Minute * 743)

	fmt.Printf("original: %v username:%v \n", a.createdAt.Hour(), a.Username)
	fmt.Printf("clone: %v username:%v", u.createdAt.Hour(), u.Username)
}
