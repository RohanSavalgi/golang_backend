package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	errorLoadingEnvFile := godotenv.Load()
	if errorLoadingEnvFile != nil {
		fmt.Println("Error in loading env file")
	}
	fmt.Println("::: Loaded the Env file :::")
}

func main() {
	
	url := "https://" + os.Getenv("DOMAIN") + "/oauth/token"

	payload := strings.NewReader("grant_type=client_credentials&client_id=" + os.Getenv("CLIENT_ID") + "&client_secret=" + os.Getenv("YOUR_CLIENT_SECRET") + "&audience=" + os.Getenv("YOUR_API_IDENTIFIER"))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}