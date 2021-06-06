# Go-Search Api

Build searching REST APIs with [Go](https://golang.org/) and use [MongoDB](https://www.mongodb.com/) for database.

Use [Fiber](https://github.com/gofiber/fiber) Web Framework for separating all routes, controllers, etc. And [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) for interact with the MongoDB instance.

## Features

- Mock data 50 documents in database by Post `/api/products/make`
- Get all data by Get `/api/products/get`
- Filter Title or Description data by Get `/api/products/get?s={data}`
- Sort Price data by Get `/api/products/get?sort={sort}` :
  - `asc` for Ascending order
  - `desc` for Descending order
- Select page data by Get `/api/products/get?page={int}`

## Flags

The following flags is available to configuration parameters.

- `-port` Port number for initial server
- `-mongo` MongoDB URI for connection
- `-db` Use or create Database
- ` -collection` Use or create collection

### go run:

```
go run main.go -port 8000 -mongo mongodb://root:root@localhost:27017 -db gp_search -collection products
```
