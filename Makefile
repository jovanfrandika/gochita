build-bot: 
	go build -o ./bin/bot ./cmd/bot/main.go

build-feedreader: 
	go build -o ./bin/feedreader ./cmd/feedreader/main.go

run-bot:
	./bin/bot

run-feedreader:
	./bin/feedreader

migrate-up:
	migrate -path db/migration -database ${db} -verbose up

migrate-down:
	migrate -path db/migration -database ${db} -verbose down

test:
	go test -cover -race ./...

clean-build:
	rm -f ./bin/*
