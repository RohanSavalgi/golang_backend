package pubsub

// import "application/logger"

func init() {
	CreateClient()
}

func PublishMessageFromGoRoutine() {
	newClient := GetClient()
	
	newClient.GetPublisher("testtopic")
	
	// _, err := newPublisher.PublishMessage("hello this is my first published message.")
	// if err != nil {
	// 	logger.ThrowErrorLog("Error in publishing the message")
	// }
}