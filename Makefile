build: 
	go build -o ./bin/app ./cmd/main.go

run:
	./bin/app

test:
	go test -cover -race ./...

clean-build:
	rm -f ./bin/*
