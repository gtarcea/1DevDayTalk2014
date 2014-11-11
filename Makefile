.PHONY: all test fmt docs

export DEVDAY_WEBDIR = ../../website

all: fmt test docs

run: bin

	(cd ./cmd/server; ./server)

bin:
	(cd ./cmd/server; go build server.go)

test:
	rm -rf test_data/t
	-go test -v ./...

docs:
	./makedocs.sh

fmt:
	-go fmt ./...
