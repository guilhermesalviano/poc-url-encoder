package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}, nil
}

// CreateQueue declares a queue
func (rmq *RabbitMQ) CreateQueue(queueName string) error {
	_, err := rmq.channel.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	return err
}

// Publish sends a message to a queue
func (rmq *RabbitMQ) Publish(queueName string, message []byte) error {
	return rmq.channel.Publish(
		"",         // exchange
		queueName,  // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
}

// Consume receives messages from a queue
func (rmq *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	return rmq.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}

// Close closes the RabbitMQ connection
func (rmq *RabbitMQ) Close() {
	if rmq.channel != nil {
		rmq.channel.Close()
	}
	if rmq.conn != nil {
		rmq.conn.Close()
	}
}

// Example usage
func ExampleRabbitMQUsage() {
	// Connection URI
	uri := "amqp://guest:guest@localhost:5672/"

	// Create RabbitMQ connection
	rmq, err := NewRabbitMQ(uri)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	// Create a queue
	queueName := "url_encoder_queue"
	err = rmq.CreateQueue(queueName)
	if err != nil {
		log.Fatalf("Failed to create queue: %v", err)
	}

	// Publish a message
	message := []byte("https://example.com")
	err = rmq.Publish(queueName, message)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Consume messages
	msgs, err := rmq.Consume(queueName)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	// Handle incoming messages
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received message: %s", msg.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}