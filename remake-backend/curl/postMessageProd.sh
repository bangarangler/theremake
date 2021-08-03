#!/bin/zsh

curl https://jonathandain.dev/api/contact \
  -H "Content-Type: application/json" \
  -d @./curl/postMessage.json \
  -X POST \
  -v | jq
