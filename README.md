# go-sdd

Todo web app by specification driven development with VS Code and Copilot. A simple web application for managing todos with authentication.

## Features

- **Signup**: POST to `/signup` - Create a new user account
- **Home Page**: GET to `/` - Welcome page with app info

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run main.go
```

The application runs on port 8081 and stores users in memory.

## Testing

Run unit tests:
```bash
go test ./...
```

Run end-to-end tests:
```bash
go test -tags=e2e ./...
```

Run all tests with coverage:
```bash
go test -cover ./...
go test -tags=e2e -cover ./...
```

## API Endpoints

### Home Page
```bash
curl http://localhost:8081/
```
Returns HTML welcome page.

### Signup
```bash
curl -X POST http://localhost:8081/signup \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

Success response:
```json
{"message": "User created"}
```

Error responses:
- Invalid JSON: `{"error": "Invalid JSON"}`
- Missing fields: `{"error": "Email and password required"}`
- Short password: `{"error": "Password must be at least 6 characters"}`
- Email exists: `{"error": "Email already exists"}`

## Project Structure

- `.sdd/`: Specifications and guidelines for development
- `main.go`: Application entry point
- `server/`: HTTP server setup
- `handlers/`: HTTP request handlers
- `*_e2e_test.go`: End-to-end tests
- `*_test.go`: Unit tests

<!--

```bash
curl -XPOST http://localhost:8080/api/todos/
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

-->

## Future Features

- User authentication (registration, login, logout)
- JWT tokens
- Route protection
- Additional todo fields (description, due date)
- User-specific todos filtering
- E2E tests
