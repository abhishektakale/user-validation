# Go API to validate user details

This is a simple Go API built using the Gin framework. The API performs validation on incoming user details requests, including custom validation for the PAN number, mobile number, and email, and logs request latency.

## Features

- **POST /submit**: Accepts a JSON payload with the following fields:
  - `name`: string
  - `pan`: string (5 letters, 4 digits, 1 letter, e.g., `ABCDE1234F`)
  - `mobile`: string (10-digit number)
  - `email`: string (valid email address)
  
- **Custom Validation**: 
  - The `pan` field is validated using a custom regular expression to ensure it matches the format.
  - The `mobile` field is validated to ensure it is exactly 10 digits long.
  - The `email` field is validated using the standard `email` validation.
  
- **Middleware**: Logs the latency of each request.

- **Dependency Injection**: The application is designed with dependency injection to ensure better testability and modularity.

## Table of Contents

- [Installation](#installation)
- [Running the API](#running-the-api)
- [Testing](#testing)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/abhishektakale/user-validation.git
   cd user-validation
   ```
2. Install the necessary Go dependencies:

   ```bash
   go mod tidy
   ```

## Running the API

1. In the project directory, run the API server:

   ```bash
   go run main.go
   ```
   This will start the server on port 8080.

## API Endpoints

### POST /user

#### Request Body (JSON)

```json
{
  "name": "John Doe",
  "pan": "ABCDE1234F",
  "mobile": "9876543210",
  "email": "john.doe@example.com"
}
```

- **name**: A string representing the name of the user (required).
- **pan**: A string representing the PAN number in the format `ABCDE1234F` (required).
- **mobile**: A string representing the 10-digit mobile number (required).
- **email**: A string representing the email address (required).

```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "pan": "ABCDE1234F",
    "mobile": "9876543210",
    "email": "john.doe@example.com"
  }'

```

#### Response

- **Success (200 OK)**:

```
{
  "message": "User created successfully"
}
```

- **Validation Error (400 Bad Request)**:

```
{
  "error": "..."
}
```

## Testing

This project includes unit tests for the handler and validation services. You can run the tests using:

```
go test -v
```

---

### Notes

- **PAN Validation**: This validation checks if the PAN number consists of 5 uppercase letters, followed by 4 digits, and ends with a single uppercase letter.
- **Mobile Number Validation**: This is checked to ensure the mobile number is exactly 10 digits.
- **Email Validation**: Uses the standard email format validation from the `validator.v10` package.