run:
	go run main.go
build:
	go build -o ./bin/convert ./cmd/convert
	chmod +x ./bin/*
link:
	go fmt ./...
clean:
	rm -rf ./bin/*