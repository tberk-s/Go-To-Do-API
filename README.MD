# TO-DO API WITH GO

The TO-DO API is a simple RESTful API developed with Go to manage tasks efficiently. This project implements core CRUD operations for managing to-do items via GET, POST, and DELETE HTTP methods.

The application supports:

- Adding tasks.
- Retrieving all tasks.
- Searching tasks by a query string.
- Deleting tasks by name.

---

## Technologies Used

- **Backend:** Go
- **Database:** PostgreSQL

## Requirements
1. Go 1.20+
1. PostgreSQL 16.0+
1. Environment variables for database configuration.

## Installation

1. Clone the repo:
    ```bash
    git clone https://github.com/tberk-s/Go-To-Do-API.git
    cd Go-To-Do-API
    ```
1. Install dependencies:
    ```bash
    go mod tidy
    ```

## Configuration

Environment variables are used to configure the database connection. Create a .env file in the project root and add the following values:

```bash
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=todo
```

## Database Setup

1. Create the db
    ```bash
    cd /path/to/Go-To-Do-API/
    source .env
    # Create a PostgreSQL user (replace 'myuser' and 'mypassword' with your username and password)
    psql -c "CREATE USER myuser WITH PASSWORD 'mypassword';"
    dropdb --if-exists "${DATABASE_NAME}"
    createdb "${DATABASE_NAME}" -O "myuser"
    ```

1. Create the table:
    ```bash
    CREATE TABLE todo_items (
    id SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    status TEXT DEFAULT 'TO_BE_STARTED'
    );
    ```

## Running the Application

1. Start the API server: `go run main.go`

2. Use Postman or command line tools such as `curl` 

| HTTP Method Endpoint | Endpoint | Description | Example Request |
|:---------------------|:------------ |:------------|:--------------|
| **GET** | /todo | Retrieve all tasks. | `curl -X GET http://localhost:8080/todo`
| **POST** | /todo | Add a new task. | `curl -X POST -d '{"item":"Buy milk"}' -H "Content-Type: application/json" http://localhost:8080/todo` 
| **GET** | /search?q={str} | Search tasks by substring query. | `curl -X GET http://localhost:8080/search?q=milk`
| **DELETE** | /todo?task={str} | Delete a task by name. | `curl -X DELETE "http://localhost:8080/todo?task=Buy%20milk"`

## Tests

It is very easy to check tests in Go with `go test ./...`