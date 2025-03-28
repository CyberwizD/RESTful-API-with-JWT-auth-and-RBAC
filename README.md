# Secure API with Role-Based Access Control (RBAC) and JWT Authentication

This repository contains a secure RESTful API built using Go that implements JWT-based authentication and Role-Based Access Control (RBAC) to manage access to protected resources. The API provides user registration, login, and access to protected routes based on user roles (e.g., Admin, User). It is designed with scalability, security, and best practices in mind, using middleware to handle authentication and authorization efficiently.

## 🚀 Key Features:
- **JWT Authentication:** Implements JSON Web Tokens (JWT) for secure authentication and token-based access.
- **Role-Based Access Control (RBAC):** Restricts access to specific routes based on user roles (Admin, User, etc.).
- **User Management:** Provides endpoints for user registration, login, and role assignment.
- **Password Hashing:** Uses `bcrypt` for secure password hashing and validation.
- **Middleware for Authentication:** Protects sensitive routes using custom JWT middleware.
- **Database Integration:** Supports PostgreSQL (via GORM) to store user data and roles.
- **Error Handling & Logging:** Includes structured logging and error handling for smooth debugging and tracing.
- **Environment Configuration:** Uses dotenv for environment-based configuration and database management.

## 🛠 Tech Stack:
- **Go (Golang):** Core language used for building the API.
- **JWT (JSON Web Token):** For stateless authentication.
- **Gorilla Mux:** HTTP router for handling API routes.
- **GORM:** ORM for interacting with PostgreSQL database.
- **bcrypt:** For password hashing and validation.
- **dotenv:** For managing environment variables.

## 🔐 Security Features:
- **Token Expiration & Validation:** Ensures tokens expire after a certain time period for enhanced security.
- **Role-Based Permissions:** Allows different levels of access based on user roles (e.g., Admins can access sensitive routes, Users cannot).
- **Password Encryption:** Ensures that user passwords are securely hashed using `bcrypt`.

## 📚 API Endpoints:
- `/register` - Register a new user.
- `/login` - Authenticate and receive a JWT token.
- `/api/profile` - Access user profile (JWT-protected route).

## 🛠 Getting Started:
1. Clone the repository:
    ```bash
    git clone https://github.com/your-username/secure-api-rbac-jwt.git
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Setup your `.env` file with the necessary environment variables (e.g., database credentials, JWT secret):
    ```ini
    DB_HOST=localhost
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=secure_api_db
    JWT_SECRET=my_secret_key
    ```

4. Run the application:
    ```bash
    go run cmd/server/main.go
    ```

## 🛡 Contributing:
Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/your-username/secure-api-rbac-jwt/issues) or open a new pull request.

## 📄 License:
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
