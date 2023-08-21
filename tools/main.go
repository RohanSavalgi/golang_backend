package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	fmt.Println("::: Loaded the Env file :::")
	
	url := "https://" + os.Getenv("DOMAIN") + "/oauth/token"
	
	payload := strings.NewReader("grant_type=client_credentials&client_id=" + os.Getenv("CLIENT_ID") + "&client_secret=" + os.Getenv("YOUR_CLIENT_SECRET") + "&audience=" + os.Getenv("YOUR_API_IDENTIFIER"))
	fmt.Println(payload)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))
}