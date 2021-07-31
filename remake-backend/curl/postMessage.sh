#!/bin/zsh

curl -X POST localhost:5000/contact \
  -H "Content-Type: application/json" \
  -d @./curl/postMessage.json \
  -v | jq
