package main

import "fmt"

type DB interface {
	Init() string
}

type App struct {
	db DB
}

func (a *App) SetDatabase(db DB) {
	a.db = db
}

func (a *App) StartWithPostgres() {
	p := &Postgres{}

	a.SetDatabase(p)
	fmt.Println(a.db.Init())
}

func (a *App) StartWithMongo() {
	m := &Mongo{}

	a.SetDatabase(m)

	fmt.Println(a.db.Init())
}

func main() {
	a := &App{}
	a.StartWithPostgres()
}
