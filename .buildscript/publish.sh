#!/bin/bash

set -e

if ! [ -x "$(command -v github-release)" ]; then
  go get -u github.com/aktau/github-release
fi

user=segmentio
repo=library-e2e-tester

version=$(shell git describe --tags --always --dirty="-dev")

# set --pre-release if there's a `-` in the tag
if [[ $version == *"-"* ]]; then
  github_release_flags := "--pre-release"
fi

github-release release \
	--security-token $GH_LOGIN \
	--user $user \
	--repo $repo \
  $(github_release_flags) \
	--tag $version \
	--name $version

for file in build/*; do
    github-release upload \
    	--security-token $GH_LOGIN \
        --user $user \
        --repo $repo \
        --tag $version \
        --name $(basename "$file") \
        --file $file
done
