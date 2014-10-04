.PHONY: all test fmt docs

all: fmt test docs

bin:
	(cd ./cmd/server; godep go build server.go)

test:
	rm -rf test_data/t
	-godep go test -v ./...

docs:
	./makedocs.sh

fmt:
	-go fmt ./...
