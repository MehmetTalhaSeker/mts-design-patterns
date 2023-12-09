package main

import (
	"fmt"
	"github.com/MehmetTalhaSeker/design-patterns/creational/builder/user"
)

func main() {
	u := user.NewBuilder().
		SetEmail("lala@lala.com").
		SetUsername("lala").
		SetPassword("123456").
		Build()

	fmt.Printf("%+v", u)
}
