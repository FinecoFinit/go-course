run:
	go run main.go
build:
	go build -o ./bin/convert ./cmd/convert
	go build -o ./bin/weeklyCalc ./cmd/weeklyCalc
	chmod +x ./bin/*
link:
	go fmt ./...
clean:
	rm -rf ./bin/*