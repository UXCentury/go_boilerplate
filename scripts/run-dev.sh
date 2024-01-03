#!/bin/bash
# Stop execution if any step fails
set -e

# Load environment variables from .env file
set -o allexport
source "./.env" set
set +o allexport

# Navigate to the root directory of the project
cd "$(dirname "$0")/.."
# Add other necessary environment variables here

# Optionally start any services (like a database) using Docker
# docker-compose -f ./path/to/docker-compose.yml up -d

# Build the application
echo "Building the application..."
go build -o ./build/project_name ./cmd/server

# Run the application
echo "Starting the application..."
./build/project_name
