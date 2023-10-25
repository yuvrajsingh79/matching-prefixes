#!/bin/bash

# Check if Go is installed
if ! [ -x "$(command -v go)" ]; then
  echo "Error: Go is not installed. Please install Go."
  exit 1
fi

# Check if Docker is installed
if ! [ -x "$(command -v docker)" ]; then
  echo "Error: Docker is not installed. Please install Docker."
  exit 1
fi

# Build the Go application
go build -o matching-prefixes ./cmd

# Check if the cache-server binary exists, and if not, build it
if [ ! -f cache-server ]; then
  echo "Building the cache-server..."
  go build -o cache-server ./internal/app
fi

# Check if the Dockerfile is present
if [ ! -f Dockerfile ]; then
  echo "Error: Dockerfile not found. Please make sure the Dockerfile is in the project directory."
  exit 1
fi

# Make the Docker container repository name lowercase
sed -i '' 's/FROM golang/FROM golang/' Dockerfile

# Build the Docker container
docker build -t matching-prefixes .

# Start the Docker container
docker run -d --name matching-prefixes -p 8080:8080 matching-prefixes

# Setup completed
echo "Setup completed. Matching-Prefixes application is running in a Docker container."
