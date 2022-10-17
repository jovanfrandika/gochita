build-bot: 
	go build -o ./bin/bot ./cmd/bot/main.go

build-livechart: 
	go build -o ./bin/livechart ./cmd/livechart/main.go

run-bot:
	./bin/bot

run-livechart:
	./bin/livechart

migrate-up:
	migrate -path db/migration -database ${db} -verbose up

migrate-down:
	migrate -path db/migration -database ${db} -verbose down

test:
	go test -cover -race ./...

clean-build:
	rm -f ./bin/*
