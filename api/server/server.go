package server

import (
	"log"
	"os"
	"time"

	"application/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	errorLoadingEnvFile := godotenv.Load()
	if errorLoadingEnvFile != nil {
		log.Println(errorLoadingEnvFile)
		log.Fatal("The env file was not loaded properly")
	}
	log.Println("Env file was loaded.")
}

func InitServer() *gin.Engine {
	// logger.InitLogger()
	logger.ThrowDebugLog("Shared : Initializing server")

	ginInstance := gin.Default()

	return ginInstance
}

func Listen(ginInstance *gin.Engine) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		logger.ThrowErrorLog("There is no SERVER_POST specified")
	}
	ginInstance.Run(":" + port)
}

func AddCors(ginInstance *gin.Engine, config *cors.Config) {
	corsConfig := config

	if corsConfig == nil {
		corsConfig = &cors.Config{
			AllowAllOrigins: true,
			AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
			AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Content-Lenght", "Accept-Encoding", "X-CSRF-Token", "Authorization", "ResponseType"},
			AllowCredentials: true,
			
			MaxAge: 12 * time.Hour,
		}
	}

	corsObject := cors.New(*corsConfig)

	ginInstance.Use(corsObject)

	logger.ThrowDebugLog("Shared: CORS added")

}