# Build stage
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and cache Go dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o event_management_service

# Runtime stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/event_management_service .

# Copy the environment variables file to the container
COPY .env .

# Set the environment variable to specify the location of the variables file
ENV ENV_FILE_PATH=.env

# Adjust permissions for the environment variables file
RUN chmod 644 .env

EXPOSE 8090

# Run the application
CMD ["./event_management_service"]
