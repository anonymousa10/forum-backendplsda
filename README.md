# Forum Backend Application

This is a simple forum backend application written in Go using the Gin framework and GORM for database operations. The application provides RESTful APIs for managing users, threads, comments, and categories.

## Setup Instructions

Follow these steps to set up and run the forum backend application on your local machine.

### Prerequisites

- Go installed on your machine
- MySQL database server

### Steps

1. **Clone the Repository:**

   git clone https://github.com/anonymousa10/forum-backendplsda.git
   cd forum-backendplsda

### Configure CORS (for Frontend Development):

The application is configured to allow requests from http://localhost:5173. Update the allowed origins in main.go if needed.

### Initialize the Database:

Ensure your MySQL server is running.

Open the models/db.go file and update the dsn variable with your MySQL connection details.

Run the following command to initialize the database:

```bash
go run main.go

This will create the necessary tables in the database.

### Run the Application:

```bash
go run main.go

The application will start on http://localhost:8080.
