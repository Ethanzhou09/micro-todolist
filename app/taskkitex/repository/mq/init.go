package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"todolist/config"
)

var RabbitMq *amqp.Connection

func InitRabbitMq() (err error) {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s", config.RabbitMQ, config.RabbitMQUser, config.RabbitMQPassWord, config.RabbitMQHost, config.RabbitMQPort)
	fmt.Println(connString)
	conn, err := amqp.Dial(connString)
	if err != nil {
		return err
	}
	RabbitMq = conn
	return
}
