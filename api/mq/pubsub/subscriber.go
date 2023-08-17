package pubsub

/*
subscriber needs 3 functions
- get sub
- get ctx
- receive the messg

struct has 2 elements
- ctx
- subcription
*/

import(
	"fmt"
	"context"

	"cloud.google.com/go/pubsub"
)

type Subscriber interface {
	GetSub() *pubsub.Subscription
	GetCtx() context.Context
	ReceiveMessage(callback func(ctx context.Context, msg *pubsub.Message)) error
}

type subscriber struct {
	ctx context.Context
	sub *pubsub.Subscription
}

func (s *subscriber) GetSub() *pubsub.Subscription {
	return s.sub
}

func (s *subscriber) GetCtx() context.Context {
	return s.ctx
}

func (s *subscriber) ReceiveMessage(callback func(ctx context.Context, msg *pubsub.Message)) error {
	err := s.sub.Receive(s.ctx, callback)

	if err != nil {
		return fmt.Errorf("Got error in Receive: %v", err)
	}

	return nil
}