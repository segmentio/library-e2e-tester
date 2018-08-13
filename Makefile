fixtures:
	go-bindata fixtures/...

build:
	gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" -parallel=4 ./...

deps:
	$Qdep ensure

vet:
	$Qgo vet -composites=false ./...

test: vet
	$Qgo test -v -cover -race ./...

.PHONY: fixtures build deps vet test
