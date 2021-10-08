#!/bin/sh

curl -X POST https://jonathandain.dev/api/contact \
  -H "Content-Type: application/json" \
  -d @./curl/automatedEmail.json \
  -v | jq
