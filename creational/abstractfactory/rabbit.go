package main

import "fmt"

type rabbitMQFactory struct{}

func (kf rabbitMQFactory) CreatePublisher() Publisher {
	return &rabbitMQPublisher{}
}

func (kf rabbitMQFactory) CreateSubscriber() Subscriber {
	return &rabbitMQSubscriber{}
}

type rabbitMQPublisher struct{}

func (kp rabbitMQPublisher) Publish(message string) {
	fmt.Printf("Producing message to RabbitMQ: %s\n", message)
}

type rabbitMQSubscriber struct{}

func (ks rabbitMQSubscriber) Subscribe(topic string) {
	fmt.Printf("Subscribing topic to RabbitMQ: %s\n", topic)
}
