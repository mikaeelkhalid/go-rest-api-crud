# Go REST API CRUD

This repository contains a sample application for a RESTful API service using Go programming language. The API allows basic CRUD operations - Create, Read, Update, Delete - on a fictional data model.

## Project Structure

```
.
├── controllers
│   ├── postController.go
├── initializers
│   ├── LoadEnv.go
|   ├── dbInit.go
├── models
│   ├── postModel.go
├── .env.sample.go
├── .gitattributes
├── .gitignore
├── README.md
└── go.mod
└── go.sum
└── main.go
```

- `controllers`: Contains controller for post.
- `initializers`: This directory contains the database connection and load envs.
- `models`: This directory contains the data model.
- `main.go`: This is the entry point of the service.

## Prerequisites

- Go (version 1.17 or later)
- A relational database system (PostgreSQL)

## Getting Started

1. Clone this repository:

```bash
git clone https://github.com/mikaeelkhalid/go-rest-api-crud.git
cd go-rest-api-crud
```

2. Install the required dependencies:

```bash
go mod tidy
```

3. Set up the environment variables for the database connection and PORT:

```bash
PORT=<port>
DB_STRING=<your_database_url>
```

4. Run the service:

```bash
go run main.go
```

The server will start on `localhost:{PORT}`.

## API Endpoints

- `POST /create`: Create a new post.
- `GET /posts`: Get all posts.
- `GET /post/{id}`: Get a post by id.
- `PUT /post/{id}`: Update a post by id.
- `DELETE /post/{id}`: Delete a post by id.

## Contribution

Feel free to contribute to this project. Any bug fixes, features, or improvements are welcome. Please create an issue before sending a PR.

## License

This project is licensed under the terms of the MIT license. For more information, see [LICENSE](LICENSE).
