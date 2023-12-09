package main

import (
	"fmt"
	"github.com/MehmetTalhaSeker/design-patterns/structural/adapter/postgres"
	"github.com/MehmetTalhaSeker/design-patterns/structural/adapter/user"
)

func main() {
	s := postgres.Store{}
	u := user.User{
		ID:       "1",
		Username: "aaa",
		Email:    "bbb@bb.com",
	}

	id, err := user.CreateUser(&s, u)
	if err != nil {
		return
	}

	fmt.Println(id)
}
