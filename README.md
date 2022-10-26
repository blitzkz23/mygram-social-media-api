# MyGram Social Media API

MyGram is a simple social media restful API, where users can post photos, comment on existing photos, and add social media to their account's profile. Created with Golang, Gin-gonic, GORM, and PostgreSQL utilizing DDD Hexagonal pattern.

## Swaggo Docs:

http://localhost:8080/swagger/index.html#/ (Still local)

## Tech Stack

- Language: Golang 1.19
- Rest Handler: Gin-gonic
- Database: PostgreSQL 14
- ORM: Gorm
- Middlewares: Authentication & Authorization with JWT Token

## Try Out Application

```bash
  # clone this repo
  $ git clone https://github.com/blitzkz23/mygram-social-media-api.git

  # download required dependencies
  $ go mod tidy

  # database configuration
  $ create new postgre db, and configure settings on database/db.go

  # run app
  $ go run main.go || make server (if you have nodemon installed)
```
