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
	// conn, err := amqp.Dial("amqp://rabbit:D66z3qm3ynC3@35.186.149.9")
	conn, err := amqp.Dial(host)
	// conn, err := amqp.Dial("amqp://rabbit:password@localhost")
	defer conn.Close()
	failOnError(err, "Fail to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"queue1", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello world 123"
	err = ch.Publish(
		"",     //exchange
		q.Name, //routing key
		false,  //mandatory
		false,  //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf("Send a msg: %s", body)
}
