package main

import "fmt"

type brokerName string

const (
	Kafka    brokerName = "kafka"
	RabbitMQ brokerName = "rabbitMQ"
)

type MessageBrokerFactory interface {
	CreatePublisher() Publisher
	CreateSubscriber() Subscriber
}

func GetMessageBrokerFactory(name brokerName) (MessageBrokerFactory, error) {
	if name == Kafka {
		return &kafkaFactory{}, nil
	}

	if name == RabbitMQ {
		return &rabbitMQFactory{}, nil
	}

	return nil, fmt.Errorf("type doesnt exists")
}
