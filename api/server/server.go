package server

import(
	"log"

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