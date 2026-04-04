# Specifications for go-sdd Application

## Implemented Features

### Signup
- **Endpoint**: POST `/signup`
- **Request Body**: `{"email": string, "password": string}`
- **Validation**:
  - Email and password required
  - Password minimum 6 characters
  - Email must be unique
- **Success Response**: `{"message": "User created"}`
- **Error Responses**: JSON with error message
- **Storage**: In-memory map (email -> password)

### Home Page
- **Endpoint**: GET `/`
- **Response**: HTML welcome page with "Welcome" heading

## Planned Features

### Authentication and Todo List

#### User Story
As a user, I want to sign up, sign in, and view my todo list after authentication. If not authenticated, I see the home page with app info and links to signup/signin.

#### Acceptance Criteria
- **Home Page (/)**:
  - If not authenticated, show welcome page with app description and links to /signup and /signin.
  - If authenticated (via JWT in header or session), redirect to /todos or show todos.
- **Signup (/signup)**: Already implemented.
- **Signin (/signin)**: POST with email/password, return JWT token.
- **Todo List (/todos)**:
  - GET: Show user's todos if authenticated
  - POST: Create new todo
  - PUT: Update todo
  - DELETE: Delete todo
- Pre-registered user: `test@go-sdd.example.com` with password `password123`, has sample todos.
- Content-Type: `application/json` for API, `text/html` for home.

#### API Specification
- **Signin (/signin)**:
  - POST: `{"email": string, "password": string}`
  - Success (200): `{"token": "jwt_token"}`
  - Error (401): `{"error": "Invalid credentials"}`
- **Todos (/todos)**:
  - GET: Headers: `Authorization: Bearer <token>`
  - Success (200): `[{"id": 1, "title": "Sample todo", "done": false}]`
  - Error (401): `{"error": "Unauthorized"}`
  - POST: Create todo
  - PUT /todos/:id: Update todo
  - DELETE /todos/:id: Delete todo

#### Notes
- Todos stored in memory per user.
- JWT token validation.
- Follow standards from `guidelines.md`.

#### Future Enhancements
- Real JWT with proper signing.
- Persistent database storage.
- User sessions.
- Todo categories/tags.