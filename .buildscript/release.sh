#!/bin/bash

set -e

if ! [ -x "$(command -v github-release)" ]; then
  go get -u github.com/aktau/github-release
fi

if ! [ -x "$(command -v gox)" ]; then
  go get github.com/mitchellh/gox
fi

make dist

user=segmentio
repo=library-e2e-tester

github-release release \
	--security-token $GH_LOGIN \
	--user $user \
	--repo $repo \
	--tag $VERSION \
	--name $VERSION

for file in build/*; do
    github-release upload \
    	--security-token $GH_LOGIN \
        --user $user \
        --repo $repo \
        --tag $VERSION \
        --name $(basename "$file") \
        --file $file
done
