package v1

import (
	"devroomtodo/storage"
	"devroomtodo/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// registerTodosRoutes registers all the routes related to Todos under the provided router group.
func registerTodosRoutes(rg *gin.RouterGroup) {
	todos := rg.Group("/todos")

	todos.GET("", getTodos)
	todos.POST("", createTodo)
	todos.GET("/:id", getTodoById)
	todos.PUT("/:id", updateTodoById)
	todos.DELETE("/:id", deleteTodoById)
}

// getTodos handles GET requests to /v1/todos
// It returns all the todos stored in the application.
func getTodos(c *gin.Context) {
	allTodos := storage.GetTodos()
	c.JSON(http.StatusOK, allTodos)
}

// createTodo handles POST requests to /v1/todos
// It expects a JSON body with title and description to create a new Todo.
func createTodo(c *gin.Context) {
	var json structs.CreateTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo := structs.Todo{
		ID:          storage.NewTodoID(),
		Title:       json.Title,
		Description: json.Description,
	}

	storage.AddTodo(newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// getTodoById handles GET requests to /v1/todos/:id
// It returns a specific Todo based on the provided ID.
func getTodoById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "ID is required")
		return
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be a number")
		return
	}

	todo := storage.GetTodoByID(idNumber)
	if todo == nil {
		c.JSON(http.StatusNotFound, "Todo not found")
		return
	}

	c.JSON(http.StatusOK, todo)
}

// updateTodoById handles PUT requests to /v1/todos/:id
// It expects a JSON body with fields to be updated in the specified Todo.
func updateTodoById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "ID is required")
		return
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be a number")
		return
	}

	var json structs.UpdateTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if storage.UpdateTodoByID(idNumber, json) {
		c.JSON(http.StatusOK, "Todo updated")
		return
	}

	c.JSON(http.StatusNotFound, "Todo not found")
}

// deleteTodoById handles DELETE requests to /v1/todos/:id
// It deletes a specific Todo based on the provided ID.
func deleteTodoById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "ID is required")
		return
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ID must be a number")
		return
	}

	if storage.DeleteTodoByID(idNumber) {
		c.JSON(http.StatusOK, "Todo deleted")
		return
	}

	c.JSON(http.StatusNotFound, "Todo not found")
}
