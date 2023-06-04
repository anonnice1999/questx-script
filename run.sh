#!/bin/bash

source .env
go run cmd/*.go generate-template -t ${ACCESS_TOKEN}
