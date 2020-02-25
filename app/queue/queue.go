package queue

import (
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/message"
	"github.com/koushamad/goro-queue/app/queue/channel"
	"github.com/koushamad/goro-queue/app/queue/rabbitmq"
)

const RABBIT  = "rabbit"
const CHANNEL = "channel"

var (
	Driver string
)

func Boot() {
	Driver = conf.Get("queue.driver")

	channel.Boot()
	switch Driver {
	case RABBIT:
		rabbitmq.Boot()
	}
}

func Push(m message.Message) {
	switch Driver {
	case RABBIT:
		rabbitmq.Push(m)
		break
	default:
		channel.Push(m)
	}
}
