package channel

import (
	"github.com/koushamad/goro-app/app/event"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/message"
	"sync"
)

type Chanel struct{}

var (
	queue *chan message.Message
	once  sync.Once
)

func Boot() {
	once.Do(func() {
		workers := conf.GetInt("queue.workers")
		Queue := make(chan message.Message)
		queue = &Queue
		for i := 1; i <= workers; i++ {
			go listener()
		}
	})
}

func Kill() {
	defer close(*queue)
}

func listener() {
	for {
		message := <-*queue
		event.Handler(message)
	}
}

func Push(m message.Message) {
	*queue <- m
}
