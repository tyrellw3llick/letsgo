# Let's Go - Go Web Application

A Go web application built while studying Let's Go by Alex Edwards. This project implements a simple, production-ready web app with features like routing, templating, and database integration, serving as a learning exercise in Go web development.

## Project Overview

The goal of this project is to follow along with the book and create a functional web application. The app includes:

- A simple HTTP server with custom routing
- HTML templating for dynamic content
- Basic CRUD operations with a database (e.g., SQLite or MySQL)
- Middleware for authentication and request handling
- Error handling and logging

## Prerequisites

- Go (version 1.21 or later recommended)
- A database (e.g., SQLite, MySQL)
- Git for version control

## Setup

1. Clone the repository:

 ```bash
 git clone https://github.com/yourusername/lets-go.git
 ```

2. Navigate to the project directory:

 ```bash
 cd lets-go
 ```

3. Install dependencies:

 ```bash
 go mod tidy
 ```

4. Run the application:

- Use flags as env virables:
- -p: port number (default is 4000)
- -u: database username:password

 ```bash
 go run ./cmd/web -p :3000 -u username:password
 ```

5. Open your browser and visit <http://localhost:4000> (or the port specified in the command).

## Progress

This repo will evolve as I work through the chapters of Let's Go. Each commit will reflect new features or improvements based on the book's lessons.

## Notes

This is a learning project, so the code may not be optimized for production use beyond what the book covers, which is in fact a lot.
