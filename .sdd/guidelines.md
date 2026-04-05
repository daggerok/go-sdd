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

- `main.go`: Entry point, sets up routes and starts the server.
- `server/`: Contains HTTP server setup and routing (package `server`).
- `handlers/`: Contains HTTP request handlers and business logic (package `handlers`).
- `main_e2e_test.go`: End-to-end tests that run against the running app.
- `.sdd/guidelines.md`: This file with coding guidelines.

**Important**: Never put handlers and server setup in the same file/package. Keep them separated:
- Server package: Only routing and server configuration
- Handlers package: Only request handling logic

## Application Requirements

- Simple HTTP server serving a welcome page on "/" and signup endpoint on "/signup".
- Use `net/http` for routing and serving.
- Handle errors gracefully without crashing the application.
- Tests should cover basic functionality, including status codes and headers.

## SDD (Specification Driven Development) Workflow

To ensure high-quality, requirement-driven development:
1. **Define Specs**: Write or update `.sdd/specs.md` with user stories, acceptance criteria, and API details. Commit specs first.
2. **Create Feature Branch**: Use `git checkout -b feature/<feature-name>` for each new feature.
3. **Write Tests (TDD)**: Add failing tests in `*_test.go` based on specs. Run `go test` to confirm they fail.
4. **Implement Code**: Write minimal code in `handlers/handlers.go` to pass tests. Follow guidelines.
5. **Refactor**: Improve code while keeping tests green. Run `go test` and `go vet`.
6. **Validate**: Ensure specs are met, update docs if needed.
7. **Merge**: Commit changes, merge to main via PR (if applicable), and delete branch.
8. **Iterate**: Based on feedback, update specs and repeat.

- Always reference `.sdd/specs.md` for requirements.
- If user specifies a different approach, adapt accordingly.
- Use learning points in responses to teach SDD concepts.

## Git Workflow

When committing and pushing changes, always use the following approach:
1. `git add .`
2. `git commit -m "proper descriptive message"`
3. `git fetch -pat`
4. `git rebase origin/$(git branch --show-current)`
5. `git push origin $(git branch --show-current)`

## Future Enhancements

- Add more routes and handlers as needed.
- Implement proper logging with a library like `logrus` or `zap`.
- Add middleware for CORS, authentication, etc.
- Expand tests with table-driven tests and more edge cases.
