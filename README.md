#### Booked


This is my attempt at a booking system for a small business. It is a work in progress and is not yet ready for production.

As a Django developer, I tried to write the code as a Django developer would. I'm sure there are so many anti-patterns in this code, but I'm learning as I GO!(pun intended)

I'm using the [Gin](https://gin-gonic.com/) framework for this project. I'm also using [Gorm](https://gorm.io/) as my ORM(Using SQLite for now).

Other libraries used:

- [jwt-go](https://github.com/golang-jwt/jwt)
- [bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

## Getting Started

### Commands

#### Run the server

```bash
go run main.go runserver
```

This will run the server on port 8100

#### Migrate the database

```bash
go run main.go migrate
```

#### Create a superuser

```bash
go run main.go createsuperuser
```
