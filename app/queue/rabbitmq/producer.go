package rabbitmq

import (
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-core/app/message"
	"github.com/streadway/amqp"
)

func (r Rabbit) Put(m message.Message) {
	err := Channel.Publish(
		r.Config.Exchange,
		r.Config.RoutingKey,
		r.Config.Mandatory,
		r.Config.Immediate,
		amqp.Publishing{
			ContentType: m.ContentType,
			Body:        m.Body,
		})
	throw.Fatal("Failed to publish a message", 118, err)
}

func Push(m message.Message) {
	rabbit := Init(m.Header)
	rabbit.Put(m)
}
