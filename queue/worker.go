package queue

import (
	"github.com/streadway/amqp"
	"log"
)

func (q *Queue) CreateQueue(ch *amqp.Channel, name string) amqp.Queue {
	queue, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	return queue
}

func (q *Queue) CreateJob(message string, ch *amqp.Channel, queue amqp.Queue) {
	err := ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
	}
}

func (q *Queue) StartQueue(ch *amqp.Channel, queue amqp.Queue, handler func()) {
	log.Printf("Queue [%v] has started...", queue.Name)
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			if string(d.Body) == "operate" {
				handler()
				log.Printf("Done")
			}
		}
	}()

	log.Printf(" [*] Waiting for jobs. To exit press CTRL+C")
	<-forever
}
