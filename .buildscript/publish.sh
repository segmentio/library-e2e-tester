#!/bin/bash

set -e

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
