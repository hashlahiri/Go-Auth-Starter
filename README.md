# Go-Auth-Starter
Using Go Language to generate a complete REST API module with authentication and authorization (JWT Token Implementation)

## Features

- User registration and login
- JWT-based authentication
- Role-based authorization
- RESTful API structure
- Basic example endpoints
- Easy to extend and customize

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or higher)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerization)
- [PostgreSQL](https://www.postgresql.org/download/) or any other preferred database

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/Go-Auth-Starter.git
    cd Go-Auth-Starter
    ```

2. Set up environment variables:

    Create a `.env` file in the root directory and add the necessary configuration:

    ```sh
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=yourdatabase
    JWT_SECRET=yourjwtsecret
    ```

3. Install dependencies:

    ```sh
    go mod tidy
    ```

4. Run the application:

    ```sh
    go run main.go

    go mod init go-auth-starter.app/base // registers the app and builds the exe file

    go mod tidy // imports all dependencies

    add extra needed dependencies whereever needed 

    go build

    go run .

    ```

### Usage

- **Register a new user:**

    ```sh
    POST /api/register
    ```

    Request body example:

    ```json
    {
      "username": "exampleuser",
      "password": "examplepassword"
    }
    ```

- **Login:**

    ```sh
    POST /api/login
    ```

    Request body example:

    ```json
    {
      "username": "exampleuser",
      "password": "examplepassword"
    }
    ```

- **Access protected route:**

    ```sh
    GET /api/protected
    ```

    Add the `Authorization: Bearer <token>` header obtained from login response.

### Project Structure

```plaintext
GoAuth-Starter/
├── main.go
├── handlers/
│   └── auth.go
├── middlewares/
│   └── auth.go
├── models/
│   └── user.go
├── routes/
│   └── routes.go
├── utils/
│   └── jwt.go
├── .env
├── go.mod
├── go.sum
└── README.md
