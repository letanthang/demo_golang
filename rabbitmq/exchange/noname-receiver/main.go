package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// conn, err := amqp.Dial("amqp://rabbit:D66z3qm3ynC3@35.186.149.9")
	conn, err := amqp.Dial("amqp://rdmdselh:coXx6LiSCvlm5noiccWyImGk0EPlre_C@mustang.rmq.cloudamqp.com/rdmdselh")
	failOnError(err, "Failed to connect RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Fail to create a channel")

	q, err := ch.QueueDeclare(
		"",
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Fail to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit presss ctrl+c")
	<-forever
}
