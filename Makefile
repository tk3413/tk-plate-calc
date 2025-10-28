APPNAME := tk-plate-calculator
VERSION := 0.1.0

generate:
	rm server_gen/server.gen.go
	go generate ./...

run:
	go run .

test:
	go test ./...

docker-build:
	docker build -t $(APPNAME):$(VERSION) .

docker-run:
	docker run -it --rm --name $(APPNAME) -e SLOG_LEVEL=DEBUG -p 8080:8080 $(APPNAME):$(VERSION)

.PHONY: generate run test docker-build docker-run
