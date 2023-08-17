package pubsub

import (
	"os"

	"application/logger"

	"github.com/gin-gonic/gin"
)

var (
	mainRecorder *os.File
)

func init() {

}

func RecorderMiddleware(methodType string, methodName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		recorderName := methodType + "+" + methodName
		PublishMessageFromGoRoutine(recorderName)
		// recorder()
		c.Next()
	}
}

func Recorder(message string) {
	recorderFile, err := os.OpenFile("recorder.txt", os.O_APPEND | os.O_CREATE , 0755)
	if err != nil {
		logger.ThrowErrorLog("Failed to create recorder file!")
	}

	recorderFile.WriteString(message + "\n")
}

