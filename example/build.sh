#!/bin/bash

set -e
set -x

GOOS=linux go build -o main

zip -r lambda.zip main index.js .env

# upload lambda.zip as lambda function
