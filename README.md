 ## e-Commerce Backend in Golang

### Features 
* e-commerce application that handles user authentication, product management, and order processing.
* Microservices Architecture
* Concurrency Control (written in go)
* Clustering and High Availability

#### Upcoming features
* Implement API rate limiting to prevent abuse.
* Message queues for asynchronous communication between microservices.
* Implement caching strategies to improve system performance.
* monitoring and alerting solutions for the microservices.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What you need to install the software:

- Go 1.18+
- Docker and Docker Compose (optional)
- PostgreSQL 

### Installation

Clone the repository:

```bash
git clone https://github.com/vigasdeep/eCommerceGo.git
cd eCommerceGo
```
rename configure .env and add values appropriately 
```bash
mv env.example .env 
```
Install Go dependencies:

```bash
go mod tidy
```
#### Running Locally

To run the application locally:

```bash
go run .
```

To run using Docker:

```bash
docker-compose up --build
```

## API Endpoints ([Postman Docs here](https://documenter.getpostman.com/view/4646448/2sA3kUGhHF))

The application provides several RESTful endpoints, grouped by functionality:

### User Authentication

- **POST** `/register`: Register a new user.
    - **Request Body Example**:
      ```json
      {
        "email": "example@example.com",
        "password": "password123"
      }
      ```
    - **Response Example**:
      ```json
      {
        "message": "User registered successfully."
      }
      ```

- **POST** `/login`: Login for existing users.
    - **Request Body Example**:
      ```json
      {
        "email": "example@example.com",
        "password": "password123"
      }
      ```
    - **Response Example**:
      ```json
      {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
      }
      ```

### Product Management

- **GET** `/api/products`: Retrieve all products.
    - **Response Example**:
      ```json
      [
        {
          "id": 1,
          "name": "Widget",
          "price": 19.99
        }
      ]
      ```

- **POST** `/api/products`: Create a new product.
    - **Request Body Example**:
      ```json
      {
        "name": "New Product",
        "price": 29.99
      }
      ```
    - **Response Example**:
      ```json
      {
        "id": 2,
        "name": "New Product",
        "price": 29.99
      }
      ```

- **PUT** `/api/products/{id}`: Update an existing product.
    - **Request Body Example**:
      ```json
      {
        "name": "Updated Product",
        "price": 39.99
      }
      ```
    - **Response Example**:
      ```json
      {
        "id": 1,
        "name": "Updated Product",
        "price": 39.99
      }
      ```

- **DELETE** `/api/products/{id}`: Delete a product.
    - **Response Example**:
      ```json
      {
        "message": "Product deleted successfully."
      }
      ```

### Order Management

- **POST** `/api/orders`: Create a new order.
    - **Request Body Example**:
      ```json
      {
        "userId": 1,
        "items": [
          {
            "productId": 1,
            "quantity": 2
          }
        ]
      }
      ```
    - **Response Example**:
      ```json
      {
        "id": 1,
        "userId": 1,
        "total": 39.98,
        "status": "Pending",
        "items": [
          {
            "productId": 1,
            "quantity": 2,
            "price": 19.99
          }
        ]
      }
      ```

- **GET** `/api/orders/{userId}`: Retrieve all orders for a user.
    - **Response Example**:
      ```json
      [
        {
          "id": 1,
          "userId": 1,
          "total": 39.98,
          "status": "Pending",
          "items": [
            {
              "productId": 1,
              "quantity": 2,
              "price": 19.99
            }
          ]
        }
      ]
      ```

### Order Items Management

- **POST** `/api/orders/{orderId}/items`: Add a new item to an order.
    - **Request Body Example**:
      ```json
      {
        "productId": 2,
        "quantity": 1
      }
      ```
    - **Response Example**:
      ```json
      {
        "orderId": 1,
        "itemId": 3,
        "productId": 2,
        "quantity": 1,
        "price": 9.99
      }
      ```

- **PUT** `/api/orders/items/{itemId}`: Update an existing order item.
    - **Request Body Example**:
      ```json
      {
        "quantity": 3
      }
      ```
    - **Response Example**:
      ```json
      {
        "orderId": 1,
        "itemId": 3,
        "productId": 2,
        "quantity": 3,
        "price": 9.99
      }
      ```

- **DELETE** `/api/orders/items/{itemId}`: Remove an item from an order.
    - **Response Example**:
      ```json
      {
        "message": "Order item deleted successfully."
      }
      ```

- **GET** `/api/orders/{orderId}/items`: Get all items for a specific

