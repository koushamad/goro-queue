package configProvider

import (
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-queue/config"
)

func Load() {
	conf.Add("Queue", config.Queue)
	conf.Add("Rabbit", config.RabbitMQ)
}
