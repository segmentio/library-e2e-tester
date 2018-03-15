fixtures:
	go-bindata fixtures/...

build:
	gox -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}" -parallel=4 ./...

.PHONY: fixtures build
