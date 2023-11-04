# Use a base image
FROM golang:1.21.0

# Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -o main .

# Specify the command to run when the container starts
CMD ["./main"]
