#!/usr/bin/env bash
set -e

if [[ -z "${RELEASE_VERSION}" ]]; then
  RELEASE_VERSION=v0.0.1
fi

docker run \
  --name hello-eks-${RELEASE_VERSION} \
  -p 32000:32000 \
  --rm \
  -e "PORT=32000" \
  docker.io/gyuho/hello-eks:${RELEASE_VERSION} \
  /hello-eks
