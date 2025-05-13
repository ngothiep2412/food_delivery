package pblocal

import (
	"g05-food-delivery/common"
	"g05-food-delivery/pubsub"
	"golang.org/x/net/context"
	"log"
	"sync"
)

// A pb run locally (in-mem)
// It has a queue (buffer channel) at it's core and many group of subscribers.
// Because we want to send a message with a specific topic for many subscribers in a group

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex // đảm bảo concurrence
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),             // 10000 -> chứa đc tổng lượng message tới 10,000
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message), // danh sách các subscriber(channel) theo từng topic
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data // gọi là enqueue
		log.Println("New event published:", data.String(), "with data", data.Data())
	}()

	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock() // thao tác vào map

	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c) // ps.mapChannel[topic] = val
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *pubsub.Message{c}
	}

	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					// remove element at index in chans
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		defer common.AppRecover()
		for {
			mess := <-ps.messageQueue // chờ message đc publish
			log.Println("Message dequeue", mess.String())

			if subs, ok := ps.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) { // xài go -> ko muốn bị đứng
						defer common.AppRecover()
						c <- mess // ửi message đến từng subscriber
					}(subs[i])
				}
			}
		}
	}()

	return nil
}
