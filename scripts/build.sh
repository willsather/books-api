#!/bin/bash

set -e

echo "Building..."

go build -o bin/books-api ./main.go

echo "Build complete."
