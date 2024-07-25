package storage

import "devroomtodo/structs"

var todos []structs.Todo

// InitTodosStorage initializes the todos slice with some default values.
func initTodosStorage() {
	todos = []structs.Todo{
		{
			ID:          0,
			Title:       "Program a todo list",
			Description: "Create a todo list application using Go and Gin",
		},
		{
			ID:          1,
			Title:       "Tidy my room",
			Description: "Clean and organize my room",
		},
	}
}

// AddTodo adds a new todo to the storage.
func AddTodo(todo structs.Todo) {
	todos = append(todos, todo)
}

// GetTodos returns all todos from the storage.
func GetTodos() []structs.Todo {
	return todos
}

// GetTodoByID returns a todo by its ID, or nil if not found.
func GetTodoByID(id int) *structs.Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo
		}
	}

	return nil
}

// UpdateTodoByID updates the title and description of a todo by its ID.
// Returns true if the update was successful, or false if the todo was not found.
func UpdateTodoByID(id int, updateTodo structs.UpdateTodo) bool {
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Update(updateTodo)
			return true
		}
	}

	return false
}

// NewTodoID generates a new ID for a todo, incrementing from the highest existing ID.
func NewTodoID() int {
	maxID := -1

	for _, todo := range todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	return maxID + 1
}

// DeleteTodoByID deletes a todo by its ID.
// Returns true if the deletion was successful, or false if the todo was not found.
func DeleteTodoByID(id int) bool {
	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}

	return false
}
