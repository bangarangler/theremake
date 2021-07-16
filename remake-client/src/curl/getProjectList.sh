#!/bin/sh

curl localhost:3000/api/projectlist -v | jq
