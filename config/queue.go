package config

import (
	"github.com/koushamad/goro-app/app/helper"
	"github.com/streadway/amqp"
)

var (
	// rabbit | chanel
	Queue = map[string]string{
		"driver":  helper.Env("QUEUE_DRIVER", "chanel"),
		"workers": helper.Env("QUEUE_WORKERS", ""),
		"traffic": helper.Env("QUEUE_TRAFFIC", ""),
	}

	RabbitMQ = map[string]string{
		"host": helper.Env("RABBIT_MQ_HOST", ""),
		"port": helper.Env("RABBIT_MQ_HOST", ""),
		"user": helper.Env("RABBIT_MQ_PORT", ""),
		"pass": helper.Env("RABBIT_MQ_PORT", ""),
	}
)

type RabbitMQConfig struct {
	QueueName     string     `default:"test"`  // queue-name
	Durable       bool       `default:"true"`  // durable
	AutoDelete    bool       `default:"false"` // delete when unused
	Exclusive     bool       `default:"false"` // exclusive
	NoWait        bool       `default:"false"` // no-wait
	Arguments     amqp.Table `default:"-"`     // arguments
	Exchange      string     `default:""`      // exchange
	RoutingKey    string     `default:"test"`  // routing key
	Mandatory     bool       `default:"false"` // mandatory
	Immediate     bool       `default:"false"` // immediate
	Consumer      string     `default:""`      // consumer
	AutoAck       bool       `default:"false"` // auto-ack
	NoLocal       bool       `default:"false"` // auto-ack
	PrefetchCount int        `default:"1"`     // prefetch-count
	PrefetchSize  int        `default:"0"`     // prefetch-count
	Global        bool       `default:"false"` // global
}
