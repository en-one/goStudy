package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOnReceiveError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}

func main() {
	//链接服务器
	conn, err := amqp.Dial("amqp://myuser:mypass@localhost:5672/")
	failOnReceiveError(err, "Faliled to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnReceiveError(err, "Fail to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnReceiveError(err, "Failed to declare a queue")

	// 获取接收消息的Delivery通道
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnReceiveError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
