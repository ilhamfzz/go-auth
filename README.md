# Golang Authentication

## Description
Implementing authentication in Golang withouth using jwt. The idea is to use a token that is generated with the user's data and then validate it with the data in the database using bearer authorization. This code using Fiber framework and PGSQl database.

## How to test
1. Clone this repository
2. Create a database in PostgreSQL with the name "go-auth"
3. Create a .env file and fill it with the following data:
```
SERVER_HOST=localhost
SERVER_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=<your password>
DB_NAME=go-auth
```
4. Run the following commands:
```
go mod tidy
go run main.go
```

## Features
### Register
#### Endpoint
```
http://localhost:8080/register
```
#### Method
```
POST
```
#### Body JSON
```
{
    "full_name": "<VARCHAR(256)>",
    "phone": "<VARCHAR(256)>",
    "username": "<VARCHAR(256)>",
    "password": "<VARCHAR(256)>"
}
```
### Generate Token
This will return a `token` that you can use to validate the user.
#### Endpoint
```
http://localhost:8080/token/generate
```
#### Method
```
POST
```
#### Body JSON
```
{
    "username": "<your username>",
    "password": "<your password>"
}
```

### Validate Token
#### Endpoint
```
http://localhost:8080/token/validate
```
#### Method
```
GET
```
#### Authorization
```
Bearer <token>
```