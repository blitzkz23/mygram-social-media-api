# MyGram Social Media API

MyGram is a simple social media restful API, where users can post photos, comment on existing photos, and add social media to their account's profile.

## Swaggo Docs:

https://mygram-social-media-api-production.up.railway.app/swagger/index.html

## Tech Stack

- Language: Golang 1.19
- Rest Handler: Gin-gonic
- Database: PostgreSQL 14
- ORM: Gorm
- Middlewares: Authentication & Authorization with JWT Token
- Testing: Testify

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
