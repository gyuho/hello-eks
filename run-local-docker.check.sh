#!/usr/bin/env bash
set -e

ENDPOINT=http://localhost:32000
echo ENDPOINT: ${ENDPOINT}

curl -L ${ENDPOINT}/hello-eks
