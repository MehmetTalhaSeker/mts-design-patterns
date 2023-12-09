package main

import "fmt"

type RabbitMQFactory struct{}

func NewRabbitFactory() *RabbitMQFactory {
	return &RabbitMQFactory{}
}

func (kf RabbitMQFactory) CreatePublisher() Publisher {
	return &rabbitMQPublisher{}
}

func (kf RabbitMQFactory) CreateSubscriber() Subscriber {
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
