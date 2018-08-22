#!/usr/bin/env bash
set -e

if [[ -z "${RELEASE_VERSION}" ]]; then
  RELEASE_VERSION=v0.0.1
fi

CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -v -o ./hello-eks

docker build \
  --tag docker.io/gyuho/hello-eks:${RELEASE_VERSION} \
  --file ./Dockerfile .

docker push docker.io/gyuho/hello-eks:${RELEASE_VERSION}
