package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://rdmdselh:coXx6LiSCvlm5noiccWyImGk0EPlre_C@mustang.rmq.cloudamqp.com/rdmdselh")
	failOnError(err, "Failed to connect RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Fail to create a channel")

	q, err := ch.QueueDeclare(
		"q0",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Fail to declare queue")

	err = ch.QueueBind(
		q.Name,      // queue name
		"",          // routing key
		"exchange1", // exchange
		false,
		nil,
	)

	failOnError(err, "Fail to bind queue")

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
