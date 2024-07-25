package tests

import (
	"devroomtodo/routes"
	"devroomtodo/storage"
	"devroomtodo/structs"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGetTodos tests the GET /v1/todos endpoint
func TestGetTodos(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/todos", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	jsonString := w.Body.String()
	todos := storage.GetTodos()
	todosBytes, _ := json.Marshal(todos)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, jsonString, string(todosBytes))
}

// TestCreateTodo tests the POST /v1/todos endpoint for creating a new todo
func TestCreateTodo(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	newTodo := structs.CreateTodo{
		Title:       "Test Todo",
		Description: "This is a test todo",
	}

	newTodoBytes, _ := json.Marshal(newTodo)
	body := strings.NewReader(string(newTodoBytes))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/todos", body)
	server.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "Test Todo")
}

// TestCreateTodoValidation tests the POST /v1/todos endpoint for validation errors
func TestCreateTodoValidation(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	newTodo := structs.CreateTodo{
		Title: "Test Todo",
	}

	newTodoBytes, _ := json.Marshal(newTodo)
	body := strings.NewReader(string(newTodoBytes))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/todos", body)
	server.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "Field validation for 'Description' failed on the 'required'")
}

// TestGetTodoById tests the GET /v1/todos/:id endpoint for retrieving a specific todo
func TestGetTodoById(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/todos/0", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Program a todo list")
}

// TestGetTodoByIdNotFound tests the GET /v1/todos/:id endpoint for a non-existent todo
func TestGetTodoByIdNotFound(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/todos/100", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Contains(t, w.Body.String(), "Todo not found")
}

// TestUpdateTodoById tests the PUT /v1/todos/:id endpoint for updating a specific todo
func TestUpdateTodoById(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	updateTodo := structs.UpdateTodo{
		Title:       "Updated Todo",
		Description: "This is an updated todo",
	}

	updateTodoBytes, _ := json.Marshal(updateTodo)
	body := strings.NewReader(string(updateTodoBytes))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/v1/todos/0", body)
	server.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Todo updated")
}

// TestUpdateTodoByIdNotFound tests the PUT /v1/todos/:id endpoint for updating a non-existent todo
func TestUpdateTodoByIdNotFound(t *testing.T) {
	server := routes.SetupRouter()
	storage.SetupStorage()

	updateTodo := structs.UpdateTodo{
		Title:       "Updated Todo",
		Description: "This is an updated todo",
	}

	updateTodoBytes, _ := json.Marshal(updateTodo)
	body := strings.NewReader(string(updateTodoBytes))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/v1/todos/100", body)
	server.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Contains(t, w.Body.String(), "Todo not found")
}
