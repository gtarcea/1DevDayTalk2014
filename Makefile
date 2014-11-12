.PHONY: all test fmt docs

export DEVDAY_WEBDIR = ../../website
export DEVDAY_CONFIG = ../../config

all: fmt test docs

run: bin

	(cd ./cmd/server; ./server)

bin:
	(cd ./cmd/server; godep go build server.go)

test:
	rm -rf test_data/t
	-godep go test -v ./...

docs:
	./makedocs.sh

fmt:
	-godep go fmt ./...
