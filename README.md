# Pixis Platform 🚀

Pixis is a modern platform for managing conscripts, departments, duties, and services. It is designed to be a complete solution, featuring a robust backend (API) and a user-friendly front-end (coming soon!).

---

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Authentication](#authentication)
- [Testing](#testing)
- [Notes](#notes)
- [License](#license)

---

## Features ✨

- JWT-based authentication for conscripts
- CRUD operations for Conscripts, Departments, Duties, Services, and Conscript-Duties relationships
- SQLite database with Gorm ORM
- Auto-generated Swagger/OpenAPI documentation
- Modular design for easy extension
- Front-end coming soon!

## Getting Started 🏁

### Prerequisites

- Go 1.20+

### Installation

1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd pixis
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Generate Swagger docs:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   swag init
   ```

### Running the Backend

```bash
go run main.go
```

The backend will be available at `http://localhost:8080`.

## Project Structure 🗂️

- `main.go` — Entry point, route setup
- `handlers/` — Route handlers (CRUD, auth, etc.)
- `models/` — Gorm models
- `database/` — DB connection and migration
- `docs/` — Auto-generated Swagger docs
- `README.md` — This file

## API Documentation 📚

- Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) for interactive API docs.
- The API is protected by JWT authentication. See the docs for details on how to authenticate and use the endpoints.

## Authentication 🔐

- Obtain a JWT by POSTing to `/auth/login` with a conscript's username and password.
- Use the returned token in the `Authorization: Bearer <token>` header for all protected endpoints.

## Testing 🧪

- Run all tests:
  ```bash
  go test ./...
  ```
- Each handler has its own test file with isolated test databases.

## Notes

- Passwords are stored in plaintext for demonstration. **Hash passwords in production!**
- JWT secret is hardcoded for demo. Use environment variables in production.
- The front-end is under development and will be integrated soon.
