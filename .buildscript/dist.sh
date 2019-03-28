#!/bin/bash

set -e

if ! [ -x "$(command -v gox)" ]; then
  go install github.com/mitchellh/gox
fi

version=$(git describe --tags --always --dirty="-dev")

gox -ldflags="-X main.Version=$version" -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...
