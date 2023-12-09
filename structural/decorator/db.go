package main

import "fmt"

type DB interface {
	Store(string)
}

type Store struct{}

func (s *Store) Store(value string) {
	fmt.Println("Storing into db: ", value)
}
