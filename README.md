# DevRoom TODO
- Part of my application to DevRoom

## How to setup development environment
- Install golang (v1.22)
- Clone this repository
- Run "go mod tidy" to install dependencies
- Run "go run main.go" to start the server
- Server will be running on http://localhost:3332

## How to run tests
- Run "go test ./tests/" to run all tests

## How to use the API
> Base URL: http://localhost:3332

### Endpoints
- GET /v1/todos
  - Get all todos
  - Response: 200 OK
  - Response Body: 
    ```json
    [
      {
        "id": 1,
        "title": "Create a todo list",
        "description": "Create a todo list using golang"
      }
    ]
    ```

- POST /v1/todos
  - Create a new todo
  - Request Body:
    ```json
    {
      "title": "Create a todo list",
      "description": "Create a todo list using golang"
    }
    ```
  - Response: 201 Created
  - Response Body:
    ```json
    {
      "id": 1,
      "title": "Create a todo list",
      "description": "Create a todo list using golang"
    }
    ```
- GET /v1/todos/:id
    - Get a todo by id
    - Response: 200 OK
    - Response Body:
        ```json
        {
        "id": 1,
        "title": "Create a todo list",
        "description": "Create a todo list using golang"
        }
        ```
      
- PUT /v1/todos/:id
  - Update a todo by id
  - Request Body:
      ```json
      {
      "title": "Create a todo list",
      "description": "Create a todo list using golang"
      }
      ```
  - Response: 200 OK
  - Response Body:
      ```
      Todo updated
      ```
- DELETE /v1/todos/:id
  - Delete a todo by id
  - Response: 200 OK
  - Response Body:
      ```
      Todo deleted
      ```
    
## How to deploy
- Run "go build" to build the binary
- Run "./devroomtodo" to start the server
- Server will be running on http://0.0.0.0:3332

## DISCLAIMER
All the responses shown here are valid only if you send a valid request like the one shown in the examples. If it is invalid the server will return a code 400 (Bad request)# devroomtodo
