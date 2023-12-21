package rabbitmq

import (
	"context"
	"fmt"
	"log"
	"ppm3/go-aim-rest-api/configs"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQActionsI interface {
	Connect() (*amqp.Connection, error)
	Ping(conn *amqp.Connection) (bool, error)
}

type RabbitMQ struct {
	ctx    context.Context
	params *configs.RabbitMQConfig
}

func NewRabbitMQConnect(ctx context.Context, params *configs.RabbitMQConfig) RabbitMQActionsI {
	return &RabbitMQ{
		ctx:    ctx,
		params: params,
	}

}

// Connect to RabbitMQ
func (r *RabbitMQ) Connect() (*amqp.Connection, error) {
	// RabbitMQ connection URL
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		r.params.Username,
		r.params.Password,
		r.params.Host,
		r.params.Port,
	)

	// Connect to RabbitMQ
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	log.Print("[OK] Connected to RabbitMQ!")

	return conn, nil
}

// Ping RabbitMQ to validate if the service is running
func (r *RabbitMQ) Ping(conn *amqp.Connection) (bool, error) {
	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return false, fmt.Errorf("failed to create channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue
	queue, err := ch.QueueDeclare(
		"ping_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return false, fmt.Errorf("failed to declare queue: %v", err)
	}

	// Publish a message to the queue
	err = ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("ping"),
		},
	)
	if err != nil {
		return false, fmt.Errorf("failed to publish message: %v", err)
	}

	// Wait for a response
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
		return false, fmt.Errorf("failed to consume messages: %v", err)
	}

	select {
	case <-msgs:
		return true, nil
	case <-time.After(5 * time.Second):
		return false, nil
	}

}
