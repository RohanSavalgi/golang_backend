package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func FetchBearerFromApi() {
	url := "https://" + os.Getenv("APPLICATION_DOMAIN") + "/oauth/token"

	payload := strings.NewReader("grant_type=client_credentials&client_id=%24%7Baccount.clientId%7D&client_secret=" + "j1TLMzuVNDX0NizHWgxVM3O1enGwOuDg6AhL2RFObzCs_kZC3mXl0BOwfv_7U632" + "&audience=" + os.Getenv("AUTH0_USER_GET_API"))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}