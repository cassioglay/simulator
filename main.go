package main

import (
	"fmt"
	"log"

	kafkaApplication "github.com/cassioglay/simulator/application/kafka"
	kafkaInfra "github.com/cassioglay/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .nv file")
	}
}

func main() {

	msgChan := make(chan *ckafka.Message)
	consumer := kafkaInfra.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaApplication.Produce(msg)
	}

	/*
		producer := kafka.NewKafkaProducer()
		kafka.Publish("Ola", "readtest", producer)
	*/

}
