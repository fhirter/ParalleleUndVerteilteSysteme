# Start with the official Golang image to build the Go application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY server.go .

# Compile the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server server.go

# Use a minimal base image to run the compiled Go application
FROM alpine:latest

# Install necessary CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/server .

# Set the environment variables. You can override these when running the container.
ENV REQUEST_DELAY_MS=100
ENV PORT=8080

# Expose the application port
EXPOSE 8080

# Command to run the Go application
CMD ["./server"]