## Technologies

- **Golang**
- **PostgreSQL**

## Getting Started

### Prerequisites

- Go (1.20 or later)
- Docker
- Docker Compose
- PostgreSQL

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/bbarmeno/test_task2.git
   cd test_task2
   ```

## Running the Application
### To Run the Application Locally


## Using Docker Compose
To build and run the application using Docker Compose, 

run:

```bash
docker-compose up --build
``` 

and the application is live!

- I've included a Postman collection. Once imported into Postman, it allows seamless access for testing the application.

## API Endpoints

```bash
GET localhost:8080/user/{userId}}/balance
```
## Request Headers:

Content-Type: application/json

- import the postman collection and run the post request while the application is running

## Responses:

- 200 OK: Successfully processed the request.
```json
{
    "userId": 1,
    "balance": "10000.00"
}
```

- 400 Bad Request: Invalid input 
- 500 Internal Server Error

```bash
POST localhost:8080/user/{userId}/transaction
```
## Request Headers:

Source-Type: client (game, server, payment)
Content-Type: application/json
Request Body:

```json
{
    "state":"win",
    "amount": 30.5,
    "transactionId": "abcd"
}
```
- import the postman collection and run the post request while the application is running

## Responses:

- 200 OK: Successfully processed the request.
- 400 Bad Request: Invalid input 
- 500 Internal Server Error



