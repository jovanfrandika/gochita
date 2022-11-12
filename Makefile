build-bot: 
	go build -o ./bin/bot ./cmd/bot/main.go

build-feedreader: 
	go build -o ./bin/feedreader ./cmd/feedreader/main.go

migrate-up:
	docker compose --profile tools run migrate

migrate-create:
	docker compose --profile tools run create-migration ${name}

up:
	docker compose up -d 

rebuild:
	docker compose build --no-cache
