#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Error: No argument provided"
    exit 1
fi

# Print the first argument
echo "Starting creation new migration named $1!"

cd ../db/migrations

echo "In $(pwd)"

../goose/goose create "$1" sql

echo "Migration created!"