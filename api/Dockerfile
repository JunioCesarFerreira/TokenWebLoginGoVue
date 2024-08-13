# api/Dockerfile
# Use official Golang image as base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8082
EXPOSE 8082

# Run the Go application
CMD ["./main"]