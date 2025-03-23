# Personal Finance Tracker API

## Introduction

This Personal Finance Tracker API helps users manage their finances by tracking budgets, transactions, and categories. Users can sign up, log in, create and manage budgets, record transactions, and categorize expenses and income.

## API Endpoints

### Authentication
- **POST `/api/v1/auth/signup`**
  - _Description:_ Register a new user.
  - _Request Body:_
    ```json
    {
      "username": "john_doe",
      "email": "john@example.com",
      "password": "password123"
    }
    ```
  - _Response:_
    ```json
    {
      "message": "User registered successfully"
    }
    ```

- **POST `/api/v1/auth/login`**
  - _Description:_ Authenticate a user and return a token.
  - _Request Body:_
    ```json
    {
      "email": "john@example.com",
      "password": "password123"
    }
    ```
  - _Response:_
    ```json
    {
      "token": "your_jwt_token"
    }
    ```

### Budgets
- **GET `/api/v1/budgets`**
  - _Description:_ Retrieve all budgets for the authenticated user.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    [
      {
        "id": 1,
        "title": "Monthly Groceries",
        "limit": 500,
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
    ```

- **POST `/api/v1/budgets`**
  - _Description:_ Create a new budget.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "title": "Monthly Groceries",
      "limit": 500
    }
    ```
  - _Response:_
    ```json
    {
      "id": 1,
      "title": "Monthly Groceries",
      "limit": 500,
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

- **GET `/api/v1/budgets/{budget_id}`**
  - _Description:_ Retrieve details of a specific budget.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    {
      "id": 1,
      "title": "Monthly Groceries",
      "limit": 500,
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

- **PUT `/api/v1/budgets/{budget_id}`**
  - _Description:_ Update an existing budget.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "title": "Monthly Groceries",
      "limit": 600
    }
    ```
  - _Response:_
    ```json
    {
      "id": 1,
      "title": "Monthly Groceries",
      "limit": 600,
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-02-01T00:00:00Z"
    }
    ```

- **DELETE `/api/v1/budgets/{budget_id}`**
  - _Description:_ Delete a budget.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    {
      "message": "Budget deleted successfully"
    }
    ```

### Transactions
- **GET `/api/v1/transactions`**
  - _Description:_ Retrieve all transactions.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    [
      {
        "id": 1,
        "user_id": 1,
        "budget_id": 1,
        "category_id": 1,
        "amount": -50,
        "description": "Groceries",
        "date": "2023-01-02T00:00:00Z",
        "created_at": "2023-01-02T00:00:00Z",
        "updated_at": "2023-01-02T00:00:00Z"
      }
    ]
    ```

- **POST `/api/v1/transactions`**
  - _Description:_ Record a new transaction (expense or income).
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "budget_id": 1,
      "category_id": 1,
      "amount": -50,
      "description": "Groceries",
      "date": "2023-01-02T00:00:00Z"
    }
    ```
  - _Response:_
    ```json
    {
      "id": 1,
      "user_id": 1,
      "budget_id": 1,
      "category_id": 1,
      "amount": -50,
      "description": "Groceries",
      "date": "2023-01-02T00:00:00Z",
      "created_at": "2023-01-02T00:00:00Z",
      "updated_at": "2023-01-02T00:00:00Z"
    }
    ```

- **GET `/api/v1/transactions/{transaction_id}`**
  - _Description:_ Retrieve details of a specific transaction.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    {
      "id": 1,
      "user_id": 1,
      "budget_id": 1,
      "category_id": 1,
      "amount": -50,
      "description": "Groceries",
      "date": "2023-01-02T00:00:00Z",
      "created_at": "2023-01-02T00:00:00Z",
      "updated_at": "2023-01-02T00:00:00Z"
    }
    ```

- **PUT `/api/v1/transactions/{transaction_id}`**
  - _Description:_ Update a transaction record.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "amount": -60,
      "description": "Groceries and Snacks",
      "date": "2023-01-02T00:00:00Z"
    }
    ```
  - _Response:_
    ```json
    {
      "id": 1,
      "user_id": 1,
      "budget_id": 1,
      "category_id": 1,
      "amount": -60,
      "description": "Groceries and Snacks",
      "date": "2023-01-02T00:00:00Z",
      "created_at": "2023-01-02T00:00:00Z",
      "updated_at": "2023-01-03T00:00:00Z"
    }
    ```

- **DELETE `/api/v1/transactions/{transaction_id}`**
  - _Description:_ Delete a transaction.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    {
      "message": "Transaction deleted successfully"
    }
    ```

### Categories
- **GET `/api/v1/categories`**
  - _Description:_ List all categories.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    [
      {
        "id": 1,
        "name": "Groceries",
        "type": "expense",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
    ```

- **POST `/api/v1/categories`**
  - _Description:_ Create a new category.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "name": "Utilities",
      "type": "expense"
    }
    ```
  - _Response:_
    ```json
    {
      "id": 2,
      "name": "Utilities",
      "type": "expense",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

- **PUT `/api/v1/categories/{category_id}`**
  - _Description:_ Update an existing category.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "name": "Monthly Utilities",
    }

- **PUT `/api/v1/categories/{category_id}`**
  - _Description:_ Update an existing category.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Request Body:_
    ```json
    {
      "name": "Monthly Utilities",
      "type": "expense"
    }
    ```
  - _Response:_
    ```json
    {
      "id": 2,
      "name": "Monthly Utilities",
      "type": "expense",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
    ```

- **DELETE `/api/v1/categories/{category_id}`**
  - _Description:_ Delete a category.
  - _Request Header:_ `Authorization: Bearer your_jwt_token`
  - _Response:_
    ```json
    {
      "message": "Category deleted successfully"
    }
    ```

---

## Sample Requests

### Authentication

**Signup:**

Request:
```sh
curl -X POST http://localhost:8080/api/v1/auth/signup -H "Content-Type: application/json" -d '{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}'
```

Response:
```json
{
  "message": "User registered successfully"
}
```

**Login:**

Request:
```sh
curl -X POST http://localhost:8080/api/v1/auth/login -H "Content-Type: application/json" -d '{
  "email": "john@example.com",
  "password": "password123"
}'
```

Response:
```json
{
  "token": "your_jwt_token"
}
```

### Budgets

**Create Budget:**

Request:
```sh
curl -X POST http://localhost:8080/api/v1/budgets -H "Authorization: Bearer your_jwt_token" -H "Content-Type: application/json" -d '{
  "title": "Monthly Groceries",
  "limit": 500
}'
```

Response:
```json
{
  "id": 1,
  "title": "Monthly Groceries",
  "limit": 500,
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```

**Get Budgets:**

Request:
```sh
curl -X GET http://localhost:8080/api/v1/budgets -H "Authorization: Bearer your_jwt_token"
```

Response:
```json
[
  {
    "id": 1,
    "title": "Monthly Groceries",
    "limit": 500,
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  }
]
```

### Transactions

**Create Transaction:**

Request:
```sh
curl -X POST http://localhost:8080/api/v1/transactions -H "Authorization: Bearer your_jwt_token" -H "Content-Type: application/json" -d '{
  "budget_id": 1,
  "category_id": 1,
  "amount": -50,
  "description": "Groceries",
  "date": "2023-01-02T00:00:00Z"
}'
```

Response:
```json
{
  "id": 1,
  "user_id": 1,
  "budget_id": 1,
  "category_id": 1,
  "amount": -50,
  "description": "Groceries",
  "date": "2023-01-02T00:00:00Z",
  "created_at": "2023-01-02T00:00:00Z",
  "updated_at": "2023-01-02T00:00:00Z"
}
```

**Get Transactions:**

Request:
```sh
curl -X GET http://localhost:8080/api/v1/transactions -H "Authorization: Bearer your_jwt_token"
```

Response:
```json
[
  {
    "id": 1,
    "user_id": 1,
    "budget_id": 1,
    "category_id": 1,
    "amount": -50,
    "description": "Groceries",
    "date": "2023-01-02T00:00:00Z",
    "created_at": "2023-01-02T00:00:00Z",
    "updated_at": "2023-01-02T00:00:00Z"
  }
]
```

### Categories

**Create Category:**

Request:
```sh
curl -X POST http://localhost:8080/api/v1/categories -H "Authorization: Bearer your_jwt_token" -H "Content-Type: application/json" -d '{
  "name": "Utilities",
  "type": "expense"
}'
```

Response:
```json
{
  "id": 2,
  "name": "Utilities",
  "type": "expense",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}
```
