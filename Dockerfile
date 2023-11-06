# Use the official Golang image as the base image
FROM golang:1.21.3

# Install air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Set the working directory in the container
WORKDIR /build

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code to the container
COPY . .

# Build the application
RUN go build -o app .

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["air"]
