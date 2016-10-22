package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://rmqhydra:hydra@10.1.2.3:5672")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create Channel
	ch, err := conn.Channel()
	handleError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"result_masscan", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	handleError(err, "Failed to declare a queue")

	// Set Prefetch count
	err = ch.Qos(1, 0, false)
	handleError(err, "Failed to set Prefetch count")

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s", d.Body)
			time.Sleep(2 * time.Second)
			log.Printf("Done processing")
			d.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for message. To exit press CTRL+C")
	<-forever
}
