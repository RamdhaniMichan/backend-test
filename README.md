# Test Naga Exchange

A simple RESTful API for user authentication and transaction management using Go, Gin, GORM, and PostgreSQL.

## Features

- User registration & login with JWT authentication
- Transaction processing and retrieval (protected endpoints)
- Dockerized for easy deployment

## Requirements

- Docker & Docker Compose
- Go (for local development)

## Getting Started

### 1. Clone the repository

```sh
git clone https://github.com/RamdhaniMichan/backend-test.git
cd test-naga-exchange
```

### 2. Environment Variables

Edit `.env` file or use Docker Compose environment section:

```
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=backend_fullstack
DB_PORT=5432
DB_SSLMODE=disable
JWT_SECRET_KEY=your_secret_key
```

### 3. Build & Run with Docker Compose

```sh
docker-compose up --build
```

- API runs on [http://localhost:8081](http://localhost:8081)
- PostgreSQL runs on port `5433` (mapped to container's `5432`)

## API Endpoints

### Auth

- `POST /register`  
  Register a new user  
  **Body:**  
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```

- `POST /login`  
  Login and get JWT token  
  **Body:**  
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```

### Transaction (Protected, require JWT in `Authorization: Bearer <token>` header)

- `GET /transaction`  
  Get user's transactions

- `POST /transaction/process`  
  Process a transaction  
  **Body:**  
  ```json
  {
    "amount": "100000",
    "type": "deposit"
  }
  ```

## Sample Request

```sh
# Register
curl -X POST http://localhost:8081/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"yourpassword"}'

# Login
curl -X POST http://localhost:8081/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"yourpassword"}'

# Get Transactions
curl -X GET http://localhost:8081/transaction \
  -H "Authorization: Bearer <JWT_TOKEN>"

# Process Transaction
curl -X POST http://localhost:8081/transaction/process \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"amount":100000,"type":"deposit"}'
```