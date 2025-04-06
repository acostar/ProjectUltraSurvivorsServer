package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

type message struct {
	UserID 	string `json:"userID"`
	Data	string `json:"data"`
}

var messages = []message{
	{ UserID: "1", Data: "[LevelLogger] LEVEL0-----" },
	{ UserID: "2", Data: "[LevelLogger] LEVEL0-----[LevelLogger] LEVEL1-----" },
	{ UserID: "3", Data: "[LevelLogger] LEVEL0-----[LevelLogger] LEVEL1-----[LevelLogger] LEVEL2-----[LevelLogger] LEVEL3-----[LevelLogger] LEVEL4-----[LevelLogger] LEVEL5-----" },
}

func main() {
	router := gin.Default()
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessages)

	router.Run("localhost:8000")
}

func getMessages(c *gin.Context) { c.IndentedJSON(http.StatusOK, messages) }

func postMessages(c *gin.Context) {
	log.SetPrefix("postMessages: ")
	log.SetFlags(0)

	var newMessage message

	if err := c.BindJSON(&newMessage); err != nil { 
		log.Fatal(err)
		return 
	}

	messages = append(messages, newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}
