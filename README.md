# Microservices/Lamdas functions in Golang

## Steps to start the environment
1. Create .env file from .env.example file
1. From root directory type `docker-compose build`
1. From root directory type `docker-compose up -d`

## Build & Run Go Application
1. From root directory type `docker-compose exec app go build /go/src/{your-app-folder}/main.go`
1. From root directory type `docker-compose exec app go run /go/src/{your-app-folder}/main.go`