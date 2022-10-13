build-bot: 
	go build -o ./bin/bot ./cmd/bot/main.go

build-livechart: 
	go build -o ./bin/livechart ./cmd/livechart/main.go

run-bot:
	./bin/bot

run-livechart:
	./bin/livechart

test:
	go test -cover -race ./...

clean-build:
	rm -f ./bin/*
