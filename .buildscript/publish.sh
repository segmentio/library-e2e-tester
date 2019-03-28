#!/bin/bash

set -ex

if ! [ -x "$(command -v github-release)" ]; then
  go get -u github.com/aktau/github-release
fi

user=segmentio
repo=library-e2e-tester

# set --pre-release if there's a `-` in the tag
if [[ $VERSION == *"-"* ]]; then
  github_release_flags := "--pre-release"
fi

github-release release \
	--security-token $GH_LOGIN \
	--user $user \
	--repo $repo \
  $(github_release_flags) \
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
