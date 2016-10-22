package main

import (
	"log"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://rmqhydra:hydra@10.1.2.3:5672/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create Channel
	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	// Declare queue
	q, err := ch.QueueDeclare(
		"result_masscan", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	handleError(err, "Failed to declare a queue")

	// Publish message to the queue
	body := "hello"
	err = ch.Publish(
		"hydra.direct", // exchange
		q.Name,         // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			// DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	handleError(err, "Failed to publish a message")
}
