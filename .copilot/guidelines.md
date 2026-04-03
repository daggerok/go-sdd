# Copilot Coding Guidelines

This file contains instructions and best practices for GitHub Copilot to follow when assisting with this Go project.

## General Rules

- **Error Handling**: Do not use `log.Fatal` or `os.Exit` in application code. Instead, return errors and handle them appropriately at the top level (e.g., in `main`).
- **Shell Commands**: When using `rm` to remove directories, always use `rm -rf` for recursive removal to avoid errors.
- **Testing**:
  - Always import the `testing` package in test files to access the `t` object.
  - Do not use `fmt` for logging in tests. Use `t.Log`, `t.Logf`, `t.Error`, or `t.Errorf` for logging and assertions.
  - Prefer assertion libraries like `testify/assert` over manual `if` checks and `t.Errorf`. Use `assert.Equal`, `assert.NotNil`, etc., for cleaner test code.
- **Code Style**:
  - Follow standard Go conventions (gofmt, go vet).
  - Export functions only when necessary (e.g., for testing or public APIs).
  - Keep handlers in separate packages for better organization and testability.

## Project Structure

- `handlers/`: Contains HTTP handlers and related logic.
- `main.go`: Entry point, sets up routes and starts the server.
- `main_test.go`: Tests for the main package and handlers.
- `.copilot/guidelines.md`: This file with coding guidelines.

## Application Requirements

- Simple HTTP server serving a welcome page on "/".
- Use `net/http` for routing and serving.
- Handle errors gracefully without crashing the application.
- Tests should cover basic functionality, including status codes and headers.

## Future Enhancements

- Add more routes and handlers as needed.
- Implement proper logging with a library like `logrus` or `zap`.
- Add middleware for CORS, authentication, etc.
- Expand tests with table-driven tests and more edge cases.