package main

import (
	"fmt"
	"sync"
)

var (
	singleton *DB
	dbOnce    sync.Once
)

type DB struct {
	i int
}

func NewDB(i int) *DB {
	dbOnce.Do(func() {
		singleton = &DB{i: i}
	})

	return singleton
}

func main() {
	wg := sync.WaitGroup{}

	for i := 1; i < 10; i++ {
		i := i
		go func() {
			wg.Add(1)
			defer wg.Done()

			a := NewDB(i)
			fmt.Println(a)
		}()
	}

	wg.Wait()

	db1 := NewDB(2)
	db2 := NewDB(3)
	db3 := NewDB(4)

	fmt.Println(db1 == db2 && db3 == db1 && db2 == db3)
}
