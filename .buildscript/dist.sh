#!/bin/bash

set -e

if ! [ -x "$(command -v gox)" ]; then
  go get github.com/mitchellh/gox
fi

gox -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" ./...
