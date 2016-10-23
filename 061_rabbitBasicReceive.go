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

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	handleError(err, "Failed to register a consumer")

	// Create a channel so that the main thread can be blocked
	// Receive will happen in the go routine
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
