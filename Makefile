include .env

.DEFAULT_GOAL := update

docs_init:
	echo "downloading necessary packages..."
	go mod download
	echo "generating specification..."
	swag init -g cmd/api/main.go

first_init:
	echo "first initialization"
	docs_init

# Updating swagger docs and than running docker-compose
start:
	echo "starting server first time..."
	
	docs_init
	echo "build and starting docker-compose with detach..."
	docker-compose up -d --build

start_production:
	echo "starting production version"
	first_init
	docker-compose up -d --build api

update:
	echo "updating project..."
	echo "pulling changes from git"
	git pull
	echo "updating documentation..."
	docs_init
	echo "build and starting docker-compose with detach..."
	docker-compose up -d --build --no-deps api

restart:
	echo "restarting docker-compose"
	docker-compose restart

down:
	echo "stopping docker-compose"
	docker-compose down

race:
	echo "checking code for race condition"
	go build -race cmd/api/main.go

migrate:
	echo migrations

test:
	echo test
