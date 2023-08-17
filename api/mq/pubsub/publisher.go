package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

/*
publisher needs 4 functions
- get topic
- get ctx
- publish message
- publish message with att

2 fields for the struct
- ctx
- topic
*/

type Publisher interface {
	GetTopic() *pubsub.Topic
	GetCtx() context.Context
	PublishMessage(msg string) (*string, error) 
	PublishMessageWithAttr(msg string, attributes map[string]string) (*string, error)
}

type publisher struct {
	ctx context.Context
	topic *pubsub.Topic
}

func (p *publisher) GetTopic() *pubsub.Topic {
	return p.topic
}

func (p *publisher) GetCtx() context.Context {
	return p.ctx
}

func (p *publisher) PublishMessage(msg string) (*string, error) {
	return p.PublishMessageWithAttr(msg, nil)
}

func (p *publisher) PublishMessageWithAttr(msg string, attributes map[string]string) (*string, error) {
	// straight up publish the message,
	// it takes 2 things data message and the attributes
	pubResult := p.topic.Publish(p.ctx, &pubsub.Message{
		Data: []byte(msg),
		Attributes: attributes,
	})

	// after the publish it gives 2 things
	// id if it is posted
	// err if it is not posted
	id, err := pubResult.Get(p.ctx)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	fmt.Printf("Published a message; msg ID : %s\n", id)
	return &id, nil
}