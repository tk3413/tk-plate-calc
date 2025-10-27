generate:
	rm server_gen/server.gen.go
	go generate ./...

run:
	go run .

test:
	go test ./...

.PHONY: generate run
