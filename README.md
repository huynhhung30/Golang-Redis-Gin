# Golang Source Code

- I am learning Golang from scratch, and I am building this project following Clean Architecture
  /Golang-Redis-Gin
  │── cmd/ # entrypoint
  │ └── api/
  │ └── main.go
  │
  │── internal/
  │ ├── domain/ # entities + interfaces (User, ...)
  │ │ └── user.go
  │ │
  │ ├── repository/ # repository interface + implementation
  │ │ ├── user_repository.go
  │ │ └── mysql/
  │ │ └── user_repo_mysql.go
  │ │
  │ ├── service/ # use case (business logic)
  │ │ └── user_service.go
  │ │
  │ ├── controller/ # orchestrator (optional)
  │ │ └── user_controller.go
  │ │
  │ ├── handler/ # HTTP handler (Gin)
  │ │ └── user_handler.go
  │ │
  │ ├── cache/ # Redis adapter
  │ │ └── redis.go
  │ │
  │ ├── utils/ # helper (error, password hash, response)
  │ │ ├── error.go
  │ │ ├── password.go
  │ │ ├── response.go
  │ │ └── log.go
  │ │
  │ └── config/ # load ENV/config
  │ └── config.go
  │
  │── pkg/ # shared libs (logger, middleware, etc.)
  │ ├── logger/
  │ └── middleware/
  │
  │── go.mod
  │── go.sum
  │── README.md

## GitHub repository containing:

## documenting
