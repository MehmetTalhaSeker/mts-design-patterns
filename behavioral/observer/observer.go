package main

import "fmt"

type Observer interface {
	OnNotify(Event)
}

type observer struct {
	id int
}

func (o *observer) OnNotify(e Event) {
	fmt.Printf(
		"observer %d received event %s\n",
		o.id, e.Data,
	)
}
