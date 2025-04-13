# Event Booking API

A Go-powered REST API interface for user authentication, event booking and management

## Features

- Empowered by the high-performance web framework **[Gin](https://github.com/gin-gonic/gin)**
- JWT(JSON Web Tokens)-based authorization for protected routes using **[JWT](https://github.com/golang-jwt/jwt)** 
- Secure password hashing by `bcrypt` from **[crypto](https://golang.org/x/crypto)**
- RESTful architecture


## Getting Started

### Prerequisites

- Go 1.21 or later
- Git

### Development Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/RenkeHuang/event_booking_api.git
   cd event_booking_api
   ```

2. Install dependencies:
   ```bash
   go mod download
   # or `go mod tidy` if error occurs
   ```

3. Start the API server locally:
   ```bash
   go run .
   ```

The API will be available at `http://localhost:8080`

## License

MIT License
