# Event Booking API

A Go-powered REST API for event booking, management and user authentication.

## Features

- Empowered by the high-performance web framework **[Gin](https://github.com/gin-gonic/gin)**
- JWT(JSON Web Tokens)-based authorization for protected routes using **[JWT](https://github.com/golang-jwt/jwt)** 
- Secure password handling by `bcrypt` from **[crypto](https://golang.org/x/crypto)**
- RESTful architecture

## Getting Started

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

The server will start at `http://localhost:8080`

### API Endpoints

#### Authentication
- `POST http://localhost:8080/signup` - Register for a new user
- `POST http://localhost:8080/login` - Login and get JWT token

#### Events
- `GET http://localhost:8080/events` - List all events
- `GET http://localhost:8080/events/:id` - Get details of an event
- `POST http://localhost:8080/events` - Create a new event
- `PUT http://localhost:8080/events/:id` - Update an event
- `DELETE [/api/events/:id](http://localhost:8080/events/:id)` - Delete an event

#### Bookings
- `POST http://localhost:8080/events/:id/register` - Book an event
- `DELETE http://localhost:8080/events/:id/register` - Cancel a booking
- `GET http://localhost:8080/events/:id/register` - View all registrations

### Usage

All protected endpoints require a JWT token which you can obtain by registering and logging in:

```bash
# Register a new user
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"username": "user1", "password": "password123", "email": "user1@example.com"}'

# Login to get a token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "user1", "password": "password123"}'
```

Use the returned token in the authorization header, e.g. cancel the registration
```bash
curl -X DELETE http://localhost:8080/events/1/register \
  -H "Authorization: YOUR_TOKEN_HERE"
```
### Project Structure
```
├── db
│   └── db.go
├── main.go
├── middlewares
│   └── auth.go
├── models
│   ├── event.go
│   └── user.go
├── routes
│   ├── events.go
│   ├── register.go
│   ├── routes.go
│   └── users.go
└── utils
    ├── hash.go
    └── jwt.go
```

## License

MIT License
