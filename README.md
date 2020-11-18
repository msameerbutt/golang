# Microservices/Lambda functions in Golang

## Steps to start the environment
1. Create .env file from .env.example file
1. From root directory type `docker-compose build`
1. From root directory type `docker-compose up -d`

## Steps to build & Run a Go Application
1. From root directory type `docker-compose exec app go build {your-app-folder}/main.go`
1. From root directory type `docker-compose exec app {your-app-folder}/main`

## Application src/helloworld : Hello World 
1. From root directory type `docker-compose exec app go build -o /go/bin/helloworld/main.exe helloworld/main.go`
1. From root directory type `docker-compose exec app /go/bin/helloworld/main.exe`

## Application src/cards : Cards 
1. From root directory type `docker-compose exec app go build -o /go/bin/cards/main.exe cards/main.go cards/deck.go`
1. From root directory type `docker-compose exec app /go/bin/cards/main.exe`