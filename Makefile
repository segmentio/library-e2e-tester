tools:
	cd tools; GO111MODULE=on go mod download

fixtures:
	go-bindata fixtures/...

dist:
	GO111MODULE=on .buildscript/dist.sh

deps:
	GO111MODULE=on go mod download

install: deps
	GO111MODULE=on go mod tidy

vet:
	GO111MODULE=on go vet -composites=false ./...

test: vet
	GO111MODULE=on go test -v -cover -race ./...

.PHONY: tools fixtures dist deps install vet test
