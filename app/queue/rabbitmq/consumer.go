package rabbitmq

import (
	"github.com/koushamad/goro-app/app/event"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-core/app/message"
	"github.com/koushamad/goro-queue/app/queue/channel"
	"time"
)

func Listener() {
	for _, qn := range event.Listen {
		rabbit := Init(qn)
		go rabbit.Consumer()
	}
}

func (r Rabbit) Consumer() {
	err := Channel.Qos(
		r.Config.PrefetchCount,
		r.Config.PrefetchSize,
		r.Config.Global,
	)
	throw.Fatal("Failed to set QoS", 116, err)

	msgs, err := Channel.Consume(
		r.Config.QueueName,
		r.Config.Consumer,
		r.Config.AutoAck,
		r.Config.Exclusive,
		r.Config.NoLocal,
		r.Config.NoWait,
		r.Config.Arguments,
	)
	throw.Fatal("Failed to register a consumer", 117, err)

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			time.Sleep(time.Duration(conf.GetInt("queue.traffic")) * time.Microsecond)
			msg := message.Message{Header: r.Config.QueueName, ContentType: m.ContentType, Body: m.Body}
			channel.Push(msg)
			m.Ack(false)
		}
	}()
	<-forever
}
