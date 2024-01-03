#!/bin/bash
# Stop execution if any step fails
set -e

# Navigate to the root directory of the project
cd "$(dirname "$0")/.."

# Define variables for build
APP_NAME="project_name"
OUTPUT_DIR="build"

# Create the output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Build the application
echo "Building $APP_NAME..."
go build -o "$OUTPUT_DIR/$APP_NAME" ./cmd/server

echo "Build completed: $OUTPUT_DIR/$APP_NAME"
