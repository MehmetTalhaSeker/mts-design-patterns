package main

type Postgres struct {
}

func (m Postgres) Init() string {
	return "Postgres migrate done."
}
