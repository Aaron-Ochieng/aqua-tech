# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod  ./

# Download and cache Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Expose the port the application runs on (if applicable)
EXPOSE 8080

# Command to run the application
CMD ["go","run","./cmd"]
