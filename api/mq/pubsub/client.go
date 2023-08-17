package pubsub

// client is the one which connects to the gcp
// and the two childer of client is publisher and subscriber
/*
interface has 4 functions
- get client
- close
- get subscriber
- get publisher

*/

import (
	"context"
	"os"
	"sync"

	"application/logger"

	"cloud.google.com/go/pubsub"
)

var (
	GCP_PROJECT_ID string
)

var setConfig = func() {
	GCP_PROJECT_ID = os.Getenv("GCP_PROJECT_ID")
}

var lock *sync.Mutex = &sync.Mutex{}
var pubsubClient Client = nil

type Client interface {
	GetClient() *pubsub.Client
	Close()
	GetSubscriber(subId string) Subscriber
	GetPublisher(topicId string) Publisher
}

type client struct {
	ctx    context.Context
	client *pubsub.Client
}

var GetClient = func() Client {
	return pubsubClient
}

func (c *client) GetClient() *pubsub.Client {
	return c.client
}

var CreateClient = func() Client {
	setConfig()
	if pubsubClient == nil {
		lock.Lock()
		defer lock.Unlock()

		if pubsubClient == nil {
			ctx := context.Background()

			c, err := pubsub.NewClient(ctx, "rohanpubsubproject")
			if err != nil {
				pubsubClient = nil 
				logger.ThrowErrorLog("Failed to create for the project with id")
			}

			pubsubClient = &client{ctx: ctx, client: c}
		}
	}

	return pubsubClient
}

func (c *client) GetSubscriber(subId string) Subscriber {
	sub := c.client.Subscription(subId)
	if sub == nil {
		logger.ThrowErrorLog("Failed to get subscription with id")
	}
	return &subscriber{ctx: c.ctx, sub: sub}
}

func (c *client) GetPublisher(topicId string) Publisher {
	t := c.client.Topic("testtopic")
	if t == nil {
		logger.ThrowErrorLog("Failed to get topic")
	}

	return &publisher{ctx : c.ctx, topic : t}
}

func (c *client) Close() {
	if pubsubClient != nil {
		c.client.Close()
		pubsubClient = nil
	} else {
		logger.ThrowErrorLog("client is not present")
	}
}
