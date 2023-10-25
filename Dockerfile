# Use the official Golang image for the build process
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code from your project to the container
COPY . .

# Build the Go application
RUN go build -o matching-prefixes ./cmd

# Build the cache-server
RUN go build -o cache-server ./internal/app

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./matching-prefixes"]
