package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"todolist/consts"
)

func SendMessage2MQ(body []byte) (err error) {
	ch, err := RabbitMq.Channel()
	if err != nil {
		return
	}
	g, _ := ch.QueueDeclare(consts.RabbitMqTaskQueue, true, false, false, false, nil)
	err = ch.Publish("", g.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		return
	}
	fmt.Println("send message to mq success")
	return
}
