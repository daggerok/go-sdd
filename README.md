# Todo Web App

A simple web application for managing todos with basic CRUD operations.

## Features

- **Create Todos**: POST to `/api/todos`
- **Get All Todos**: GET to `/api/todos`
- **Get Todo Detail**: GET to `/api/todos/:id`
- **Update Todo**: PUT to `/api/todos/:id`
- **Delete Todo**: DELETE to `/api/todos/:id`

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run main.go
```

The application will create a SQLite database file (todos.db) automatically.

## Testing

Run all tests:
```bash
./run_tests.sh
```

Or manually:
```bash
go test -v
go test -cover
```

## API Endpoints

### Get All Todos
```bash
curl http://localhost:8080/api/todos
```

### Create Todo
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "title": "My first todo", "completed": false}'
```

### Get Todo by ID
```bash
curl http://localhost:8080/api/todos/1
```

### Update Todo
```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Updated todo", "completed": true}'
```

### Delete Todo
```bash
curl -X DELETE http://localhost:8080/api/todos/1
```

## Future Features

- User authentication (registration, login, logout)
- JWT tokens
- Route protection
- Additional todo fields (description, due date)
- User-specific todos filtering
- E2E tests