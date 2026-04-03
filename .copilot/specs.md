# Specifications for go-sdd Application

## Feature: Authentication and Todo List

### User Story
As a user, I want to sign up, sign in, and view my todo list after authentication. If not authenticated, I see the home page with app info and links to signup/signin.

### Acceptance Criteria
- **Home Page (/todos)**:
  - If authenticated (via JWT in header or session), show user's todo list (HTML or JSON).
- **Home Page (/)**:
  - If not authenticated, show welcome page with app description and links to /signup and /signin.
- **Signup (/signup)**: As above.
- **Signin (/signin)**: As above.
- **Todo List (/todos)**: As above.
- Pre-registered user: `test@go-sdd.example.com` with password `password123`, has sample todos.
- Content-Type: `application/json` for API, `text/html` for home.

### API Specification
- **Home (/)**:
  - GET: Conditional response based on auth.
- **Signup (/signup)**: As above.
- **Signin (/signin)**: As above.
- **Todos (/todos)**:
  - Method: GET
  - Headers: `Authorization: Bearer <token>`
  - Success (200): `[{"id": 1, "title": "Sample todo", "done": false}]`
  - Error (401): `{"error": "Unauthorized"}`

### Notes
- Todos stored in memory per user.
- Simple token validation (check if token == "valid_token" for test user).
- Follow standards from `guidelines.md`.

### Future Enhancements
- Real JWT.
- CRUD for todos.
- Database.