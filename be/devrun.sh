#!/bin/bash

# Change to the cmd directory
cd cmd

# Load additional environment variables from env.sh file (the standard)
if [ -f env.sh ]; then
    source env.sh
else
    echo "Error: env.sh file not found."
    exit 1
fi

# Load environment variables from .env file, ignoring comments, empty lines, and stripping quotes
if [ -f .env ]; then
    while IFS='=' read -r key value; do
        # Ignore comments and empty lines
        if [[ ! "$key" =~ ^# && -n "$key" ]]; then
            # Remove any surrounding quotes from the value
            value="${value%\"}"
            value="${value#\"}"
            export "$key=$value"
        fi
    done < .env
else
    echo "Error: .env file not found."
    exit 1
fi

# Run the Go application
go run main.go  # Adjust this command based on your entry point
