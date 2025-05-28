# JWT Token Service

This project provides a simple Go-based API for generating and extracting JSON Web Tokens (JWTs) using email addresses. It includes two main functionalities:

1. **Generate Token**: Creates a JWT token for a given email address.
2. **Extract Email**: Extracts the email address from a provided JWT token.

The API is built using the [Fiber](https://gofiber.io/) web framework and the [jwt-go](https://github.com/golang-jwt/jwt) library for JWT operations. It is designed to be lightweight, secure, and easy to integrate.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [Generate New Token](#generate-new-token)
  - [Extract Email from Token](#extract-email-from-token)
- [Testing with Postman](#testing-with-postman)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features

- Generate JWT tokens with an email address embedded in the claims.
- Extract email addresses from valid JWT tokens.
- Input validation for email format and required fields.
- Error handling with consistent JSON responses.
- Environment-based configuration for JWT secret.

## Prerequisites

- **Go**: Version 1.16 or higher.
- **Git**: For cloning the repository.
- **Postman**: (Optional) For testing the API endpoints.
- **Environment Variables**:
  - `JWT_SECRET`: A secret key for signing and verifying JWT tokens.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ajayashokgokhale/go-jwt-token.git
   cd jwt-token-service
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

   Ensure you have the following dependencies in your `go.mod`:
   - `github.com/gofiber/fiber/v2`
   - `github.com/golang-jwt/jwt/v4`

3. Set up the environment variable for the JWT secret:
   ```bash
   export JWT_SECRET="your-secure-secret-key"
   ```

   Alternatively, you can create a `.env` file and load it using a library like `godotenv` (not included in the provided code).

## Configuration

- **JWT_SECRET**: This environment variable must be set to a secure, random string used for signing and verifying JWT tokens. Do not hardcode this in the source code.
- **Port**: The application runs on port `8080` by default. Modify the Fiber configuration if you need a different port.

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

   The API will be available at `http://localhost:8080`.

2. Verify the server is running by accessing the endpoints (see [API Endpoints](#api-endpoints)).

## API Endpoints

### Generate New Token

- **Endpoint**: `POST /api/newtoken`
- **Description**: Generates a JWT token for a given email address.
- **Request Body**:
  ```json
  {
      "email": "abc@example.com"
  }
  ```
- **Response** (Success):
  ```json
  {
      "data": {
          "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
      },
      "message": "Token generated successfully",
      "success": true
  }
  ```
- **Response** (Error):
  ```json
  {
      "message": "Invalid email format",
      "success": false
  }
  ```
- **Status Codes**:
  - `200 OK`: Token generated successfully.
  - `400 Bad Request`: Invalid JSON or email format.
  - `500 Internal Server Error`: Token generation failed.

### Extract Email from Token

- **Endpoint**: `POST /api/tokenextract`
- **Description**: Extracts the email address from a provided JWT token.
- **Request Body**:
  ```json
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```
- **Response** (Success):
  ```json
  {
      "data": {
          "email": "abc@example.com"
      },
      "message": "Token extract successfully",
      "success": true
  }
  ```
- **Response** (Error):
  ```json
  {
      "message": "Invalid token: token is expired",
      "success": false
  }
  ```
- **Status Codes**:
  - `200 OK`: Email extracted successfully.
  - `400 Bad Request`: Invalid JSON or missing token.
  - `401 Unauthorized`: Invalid or expired token.
  - `500 Internal Server Error`: Token parsing failed.

## Testing with Postman

A Postman collection is included in the repository (`JWT Tokens.postman_collection.json`). Import it into Postman to test the API:

1. Open Postman and import the collection:
   - File > Import > Choose `JWT Tokens.postman_collection.json`.
2. Set the `JWT_SECRET` environment variable in your Postman environment or ensure it matches the server's configuration.
3. Use the following requests:
   - **New Token**: Send a POST request to `http://localhost:8080/api/newtoken` with the email in the body.
   - **Token Extract**: Send a POST request to `http://localhost:8080/api/tokenextract` with the token in the body.

Example requests are pre-configured in the collection.

## Project Structure

```
go-jwt-token/
├── actions
│   ├── newtoken
│   │   ├── handler.go
│   │   └── logic.go
│   └── tokenextract
│       ├── handler.go
│       └── logic.go
├── go.mod
├── go.sum
├── gtservices
│   ├── jwtgenx
│   │   ├── jwt.go
│   │   └── middleware.go
│   ├── responsex
│   │   └── response.go
│   └── utils
│       └── email.go
├── LICENSE
├── main.go
├── pkg
│   └── dbx
│       └── db.go
├── PostmanCollection
│   └── JWT Tokens.postman_collection.json
├── README.md
└── request
    ├── newtoken.go
    ├── request.go
    └── tokenextract.go

```

- **`newtoken/`**: Handles token generation logic and HTTP endpoint.
- **`tokenextract/`**: Handles email extraction from tokens and HTTP endpoint.
- **`gtservices/`**: Contains shared utilities for JWT operations, response formatting, and other helpers.
- **`main.go`**: Entry point for the Fiber application (not provided but assumed).
- **`JWT Tokens.postman_collection.json`**: Postman collection for testing.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -m "Add your feature"`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please ensure your code follows Go best practices and includes tests where applicable.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.