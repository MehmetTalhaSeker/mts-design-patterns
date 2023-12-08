package main

func main() {
	kf, err := GetMessageBrokerFactory(Kafka)
	if err != nil {
		return
	}

	kp := kf.CreatePublisher()
	ks := kf.CreateSubscriber()

	kp.Publish("HI")
	ks.Subscribe("top pic")

	rf, err := GetMessageBrokerFactory(RabbitMQ)
	if err != nil {
		return
	}

	rp := rf.CreatePublisher()
	rs := rf.CreateSubscriber()

	rp.Publish("HI")
	rs.Subscribe("top pic")
}
