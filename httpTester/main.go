// package main
// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.GET("/hello/:id", func(c *gin.Context) {
// 		id := c.Param("id")

// 		c.JSON(http.StatusOK, gin.H{ "data is " : id })
// 	})

// 	r.Run()
// }

// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// type User struct {
//     ID   int     `json:"id"`
//     Name string  `json:"name"`
//     Age  int     `json:"age"`
// 	Email string `json:"email"`
// }

// func main() {
// 	user1 := User{
// 		ID: 1,
// 		Name: "rohan",
// 		Age: 23,
// 		Email: "rohan@rohan.com",
// 	}

// 	server := gin.Default() 

// 	server.GET("/hello/:id", func(c *gin.Context){
// 		// id := c.Param("id")

// 		c.JSON(http.StatusOK, user1)
// 	})

// 	server.Run()
// }

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"encoding/json"
)

type user struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func getMethod(c *gin.Context) {
	// id := c.Param("id")
	user1 := user{
		Id: 1,
		Name: "phanish and neeraj",
		Email: "rohan@okbrother.com",
	}

	c.JSON(http.StatusOK, user1)
}

func postMethod(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err)
		return 
	}

	var inputUser user

	errUnmarsh := json.Unmarshal(body, &inputUser)
	if errUnmarsh != nil {
		c.JSON(400, errUnmarsh)
		return
	}

	c.JSON(http.StatusOK, inputUser)
}

func main() {
	server := gin.Default()


	server.GET("/hello/:id", getMethod)
	server.DELETE("/hello", postMethod)

	server.Run()
}