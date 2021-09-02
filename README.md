# Go REST API Example
A RESTful API example for simple users application with Go

This is a simple tutorial or example to create a simple RESTful API with Go using gin (The gin library is very easy to understand) and gorm (modeling database)

## Installation & Run
```bash
# Download this project
go get github.com/c0rz/go-user-restful-api
```

Before running API server, you should set the database config with yours or set the your database config with my values on [main.go](https://github.com/c0rz/go-user-restful-api/blob/main/main.go#L15)
```go
connect := "root:@tcp(127.0.0.1:3306)/golangdatabase?charset=utf8mb4&parseTime=True&loc=Local"
```
```bash
# Build and Run
cd go-user-restful-api
go build
go run main.go

# API Endpoint : http://127.0.0.1:8080
```

## API

#### /list
* `GET` : Get all users

#### /create
* `POST` : Create a new user

#### /update/:id
* `PUT` : Update user

#### /delete/:id
* `DELETE` : Delete user
