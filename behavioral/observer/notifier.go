package main

type Notifier interface {
	Register(Observer)
	Unregister(Observer)
	Notify(Event)
}

type notifier struct {
	observers map[Observer]struct{}
}

func (n *notifier) Register(o Observer) {
	n.observers[o] = struct{}{}
}

func (n *notifier) Unregister(o Observer) {
	delete(n.observers, o)
}

func (n *notifier) Notify(e Event) {
	for o := range n.observers {
		o.OnNotify(e)
	}
}
