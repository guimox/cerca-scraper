package queue

import (
	"cerca-scraper/internal/schedule"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQConfig(url string) (*RabbitMQConfig, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	err = ch.ExchangeDeclare(
		"scraper.fanout",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	return &RabbitMQConfig{
		conn:    conn,
		channel: ch,
	}, nil
}

func (r *RabbitMQConfig) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

func (r *RabbitMQConfig) PublishSchedule(data schedule.TableData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data: %v", err)
	}

	err = r.channel.Publish(
		"scraper.fanout",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		},
	)
	if err != nil {
		return fmt.Errorf("error publishing message: %v", err)
	}

	log.Printf("Published schedule data for station: %s", data.Station)
	return nil
}
