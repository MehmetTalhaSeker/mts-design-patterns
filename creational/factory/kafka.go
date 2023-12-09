package main

import "fmt"

type KafkaFactory struct{}

func NewKafkaFactory() *KafkaFactory {
	return &KafkaFactory{}
}

func (kf KafkaFactory) CreatePublisher() Publisher {
	return &kafkaPublisher{}
}

func (kf KafkaFactory) CreateSubscriber() Subscriber {
	return &kafkaSubscriber{}
}

type kafkaPublisher struct{}

func (kp kafkaPublisher) Publish(message string) {
	fmt.Printf("Producing message to Kafka: %s\n", message)
}

type kafkaSubscriber struct{}

func (ks kafkaSubscriber) Subscribe(topic string) {
	fmt.Printf("Subscribing topic to Kafka: %s\n", topic)
}
