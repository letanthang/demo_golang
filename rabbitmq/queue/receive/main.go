package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("RABBIT_HOST", "amqp://guest:guest@localhost:5672")
	host := viper.GetString("RABBIT_HOST")
	fmt.Println("RABBIT_HOST", host)
	//"amqp://rdmdselh:coXx6LiSCvlm5noiccWyImGk0EPlre_C@mustang.rmq.cloudamqp.com/rdmdselh"
	conn, err := amqp.Dial(host)
	failOnError(err, "Failed to connect RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Fail to create a channel")

	q, err := ch.QueueDeclare(
		"queue1",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Fail to declare queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for msg := range msgs {
		fmt.Println("receive message", string(msg.Body))
	}

}
