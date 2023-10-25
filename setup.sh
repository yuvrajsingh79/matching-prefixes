#!/bin/bash

# Build the Go application
echo "Building the Matching-Prefixes application..."
go build

# Start the cache server (You can replace 'cache-server' with the actual command to start your cache server)
echo "Starting the cache server..."
./cache-server &

# Run the Go application in a Docker container
echo "Building the Docker container..."
docker build -t Matching-Prefixes .

# Run the Docker container
echo "Starting the Docker container..."
docker run -d -p 8080:8080 Matching-Prefixes

echo "Setup completed. Matching-Prefixes application is running in a Docker container."
