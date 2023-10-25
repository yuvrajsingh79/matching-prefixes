# Use the official Golang image as the base image.
FROM golang:latest

# Set the working directory in the container.
WORKDIR /app

# Copy the Go application source code into the container.
COPY . .

# Build the Go application.
RUN go build -o app

# Expose the port your application will listen on.
EXPOSE 8080

# Define the command to run your application.
CMD ["./app"]
