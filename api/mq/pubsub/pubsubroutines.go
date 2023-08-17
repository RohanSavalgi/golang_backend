package pubsub

import (
	"context"

	"application/logger"

	"cloud.google.com/go/pubsub"
)

// import "application/logger"

func init() {
	CreateClient()
}

func PublishMessageFromGoRoutine(message string) {
	newClient := GetClient()

	newPublisher := newClient.GetPublisher("testtopic")
	
	_, err := newPublisher.PublishMessage(message)
	if err != nil {
		logger.ThrowErrorLog("Error in publishing the message")
	}
}

func RecieveMessageFromGoRoutine() string {
	newClient := GetClient()

	newSubscriber := newClient.GetSubscriber("testtopic-sub")
	result := ""
	newSubscriber.ReceiveMessage(func(ctx context.Context, msg *pubsub.Message) {
		defer func() {
			if r := recover(); r != nil {
				logger.ThrowErrorLog("Recovering from a panic!")
			}
		}()
		msg.Ack()
		Recorder(string(msg.Data))
	})
	return result
}