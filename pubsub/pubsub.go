package pubsub

import "golang.org/x/net/context"

type Topic string

type Pubsub interface {
	Publish(ctx context.Context, channel Topic, data *Message) error
	Subscribe(ctx context.Context, channel Topic) (ch <-chan *Message, close func()) // golang -> xài channel hay hơn callback
	// streaming -> sẽ gây ra block

	//! Tại sao ko viết hàm unsubscribed -> vì viết hàm thì cần truyền vào, khi truyền vào channel nó sẽ làm unsubscribe hết channel
	//UnSubscribed(ctx context.Context, channel Channel) error
}
