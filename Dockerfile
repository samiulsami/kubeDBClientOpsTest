# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my-app .

# Stage 2: Create the final image with the compiled binary
FROM alpine:latest

# Set working directory in the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/my-app .

# Expose port 8080 (adjust according to your application)
EXPOSE 8080

# Run the Go binary
CMD ["./my-app", "pgTestAll"]
