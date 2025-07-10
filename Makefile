run:
	go run main.go
build:
	go build -o ./bin/convert ./cmd/convert
link:
	go fmt ./...
clean:
	rm -rf ./bin/*