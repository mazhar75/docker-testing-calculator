# Use the official Go image as a builder
FROM golang:1.22.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum if you have them (for dependency caching)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app (you can name the binary 'calculator' or anything you like)
RUN go build -o calculator .

# Use a minimal image for the final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/calculator .

# Command to run the binary
CMD ["./calculator"]
