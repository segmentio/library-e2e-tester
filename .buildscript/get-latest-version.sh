#!/bin/bash

# Usage: get-latest-version.sh [executable_name]
# Download the latest release, for the given executable (default to "tester_linux_amd64")

executable=$1

if [[ -z ${executable} ]] ; then
    executable="tester_linux_amd64"
fi

wget $(curl -s https://api.github.com/repos/segmentio/library-e2e-tester/releases/latest | jq -r '.assets[].browser_download_url' | grep ${executable})
