package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"Id"`
	Item      string `json:"Item"`
	Completed bool   `json:"Completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Sleep", Completed: false},
}

func getTodos(context *gin.Context) { //Get all array endpoints
	r := len(todos)
	if r == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "No todos found"})
	}
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodo(context *gin.Context) { //Get single endpoint
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) { //Add single endpoint
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context) { //Toggle completedness
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) { //Search for specific todo by ID
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
