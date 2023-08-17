#!/bin/bash

while read -ra e; do
  export $e
done <<<"$(cat .env)"

go run cmd/*.go generate-template
