package main

func main() {
	kf := NewKafkaFactory()

	kp := kf.CreatePublisher()
	ks := kf.CreateSubscriber()

	kp.Publish("HI")
	ks.Subscribe("top pic")

	rf := NewRabbitFactory()

	rp := rf.CreatePublisher()
	rs := rf.CreateSubscriber()

	rp.Publish("HI")
	rs.Subscribe("top pic")
}
