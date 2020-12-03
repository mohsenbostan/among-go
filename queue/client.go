package queue

import (
	"github.com/mohsenbostan/among-go/utils"
	"github.com/streadway/amqp"
	"log"
	"os"
)

type Queue struct {
	Hostname string
	Username string
	Password string
	Port     string
}

func NewQueue() *Queue {
	utils.LoadEnvVariables()

	return &Queue{
		Hostname: os.Getenv("RABBITMQ_DEFAULT_HOSTNAME"),
		Username: os.Getenv("RABBITMQ_DEFAULT_USER"),
		Password: os.Getenv("RABBITMQ_DEFAULT_PASS"),
		Port:     os.Getenv("RABBITMQ_DEFAULT_PORT"),
	}
}

func (q *Queue) Client() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}

	return conn, ch
}
