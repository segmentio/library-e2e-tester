#!/bin/bash

set -e

version=$(git describe --tags --always --dirty="-dev")

gox -ldflags="-X main.Version=$version" -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...
