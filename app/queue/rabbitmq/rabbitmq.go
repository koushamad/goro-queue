package rabbitmq

import (
	"github.com/creasty/defaults"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-queue/config"
	"github.com/streadway/amqp"
	"sync"
)

type Rabbit struct {
	Queue amqp.Queue
	Config *config.RabbitMQConfig
	Messages []map[string]*chan amqp.Delivery
}

var (
	Connection *amqp.Connection
	Channel *amqp.Channel
	once sync.Once
)

func Boot()  {
	once.Do(func() {
		conn, err := amqp.Dial(getConnection())
		throw.Fatal("Cannot Connect To Rabbit", 119, err)
		Connection = conn
		ch, err := conn.Channel()
		throw.Fatal("Failed to open a channel", 120, err)
		Channel = ch
		go Listener()
	})
}

func getConnection()string{
	return "amqp://" + conf.Get("rabbit.user") +
		":" + conf.Get("rabbit.pass") +
		"@" + conf.Get("rabbit.host") +
		":" + conf.Get("rabbit.port")
}

func Kill()  {
	defer Connection.Close()
	defer Channel.Close()
}

func Init(qn string) Rabbit {
	config := GetDefaultConfig(qn)
	return connect(config)
}

func connect(c *config.RabbitMQConfig) Rabbit {
	q, err := Channel.QueueDeclare(
		c.QueueName,
		c.Durable,
		c.AutoDelete,
		c.Exclusive,
		c.NoWait,
		c.Arguments,
	)
	throw.Fatal("Failed to declare a queue", 121, err)
	return Rabbit{Config: c, Queue:q}
}

func GetDefaultConfig(qn string) *config.RabbitMQConfig {
	conf := &config.RabbitMQConfig{}
	throw.Fatal("Cannot set producer default value", 122, defaults.Set(conf))

	conf.QueueName = qn
	conf.RoutingKey = qn
	conf.Arguments = amqp.Table{}

	return conf
}