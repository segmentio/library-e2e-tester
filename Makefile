fixtures:
	go-bindata fixtures/...

build:
	gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" -parallel=4 ./...

deps:
	GO111MODULE=on go mod download

install: deps
	GO111MODULE=on go mod tidy

vet:
	GO111MODULE=on go vet -composites=false ./...

test: vet
	GO111MODULE=on go test -v -cover -race ./...

.PHONY: fixtures build deps install vet test
