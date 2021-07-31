#!/bin/zsh

curl localhost:5000/ping \
  -H "Content-Type: application/json" -v | jq

