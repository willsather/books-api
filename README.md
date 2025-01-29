# **Books API**

A simple RESTful API for managing a collection of books using **Go** (`net/http`).

## **Features**
- Get all books
- Get a specific book by ID
- Add a new book
- Delete a book by ID

---

## **Table of Contents**
1. [Setup and Installation](#setup-and-installation)
2. [Usage](#usage)
3. [How It Works](#how-it-works)

---

## **Setup and Installation**

### **Prerequisites**
- Go 1.23 or higher
- Git (to clone the repository)

### **Steps to Run the Project**
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/your-username/books-api.git
   cd books-api
   ```

2. **Install Dependencies**:
   The project uses Go Modules to manage dependencies. Run:
   ```bash
   go mod tidy
   ```

3. **Run the API Server**:
   To start the server, run:
   ```bash
   go run main.go
   ```

4. **Server Running**:
   By default, the API will be available at `http://localhost:8080`.

---

## **Usage**

You can use `curl` or save requests in Postman to interact with the API. Here's a general usage example:

1. **Get All Books** (using `curl`):
   ```bash
   curl http://localhost:8080/books
   ```

2. **Add a New Book** (using `curl`):
   ```bash
   curl -X POST -d '{"title": "Go Programming", "author": "John Doe", "isbn": "1234567890123"}' \
   http://localhost:8080/books/create
   ```

3. **Delete a Book** (using `curl`):
   ```bash
   curl -X DELETE http://localhost:8080/books/1
   ```

> Alternatively, you can use **Bruno** to save these requests for easy reuse.

---

## **How It Works**

### **Project Structure**

```
/books-api
│── go.mod                  # Go module file
│── main.go                 # Entry point, sets up routes and starts the server
│── /internals              # Contains internal logic for books
│   ├── /books
│   │   ├── handler.go      # HTTP request handlers for book routes
│   │   ├── service.go      # Business logic for handling books (CRUD)
```

### **Core Components**

- **Service Layer** (`service.go`):  
  The core logic for managing books (adding, deleting, retrieving) is handled here. It uses a map to store books and provides functions to perform CRUD operations.

- **Handler Layer** (`handler.go`):  
  This layer contains HTTP request handlers for interacting with the service layer. Each HTTP request maps to a function that performs the necessary actions (e.g., adding or removing books).

- **Main Application** (`main.go`):  
  The `main.go` file starts the web server and routes HTTP requests to the appropriate handler.
