package main

// stolen from -> https://pavelfokin.dev/blog/how-to-implement-observer-pattern-in-go/
func main() {
	smsNotifier := notifier{
		observers: map[Observer]struct{}{},
	}

	emailNotifier := notifier{
		observers: map[Observer]struct{}{},
	}

	o1 := &observer{1}
	o2 := &observer{2}

	smsNotifier.Register(o1)
	smsNotifier.Register(o2)

	smsNotifier.Notify(Event{"discount for you: iphone"})

	smsNotifier.Unregister(o1)

	smsNotifier.Notify(Event{"dont you leave me too o2"})

	smsNotifier.Unregister(o2)
	smsNotifier.Notify(Event{":("})

	emailNotifier.Register(o1)
	emailNotifier.Notify(Event{"welcome to email subscription"})
}
