package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	kafkaApp "github.com/yorigcs/simulator-delivery/application/kafka"
	"github.com/yorigcs/simulator-delivery/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApp.Produce(msg)

	}

}
