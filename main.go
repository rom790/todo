package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"iten"`
	Completed bool   `json:"completed"`
}

var todo = []Todo{
	{ID: "1", Item: "read", Completed: false},
	{ID: "2", Item: "go out", Completed: false},
	{ID: "2", Item: "sleep", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8081")
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todo)
}
