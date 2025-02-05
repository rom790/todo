package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todo = []Todo{
	{ID: "1", Item: "read", Completed: false},
	{ID: "2", Item: "go out", Completed: false},
	{ID: "3", Item: "sleep", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getAllTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:8081")
}

func getAllTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) {
	var newTodo Todo

	if err := context.Bind(&newTodo); err != nil {
		return
	}

	todo = append(todo, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}
	fmt.Println(todo)
	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*Todo, error) {
	for i, td := range todo {
		if id == td.ID {
			return &todo[i], nil
		}
	}
	return nil, errors.New("no records with this id")
}
