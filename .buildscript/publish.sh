#!/bin/bash

set -e

if ! [ -x "$(command -v github-release)" ]; then
  (cd /tmp; go get -u github.com/mitchellh/gox)
fi

user=segmentio
repo=library-e2e-tester

version=$(git describe --tags --always --dirty="-dev")

github-release release \
	--security-token $GH_LOGIN \
	--user $user \
	--repo $repo \
	--tag $version \
	--name $version

for file in dist/*; do
    github-release upload \
    	--security-token $GH_LOGIN \
        --user $user \
        --repo $repo \
        --tag $version \
        --name $(basename "$file") \
        --file $file
done
