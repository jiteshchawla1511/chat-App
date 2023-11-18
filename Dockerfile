
# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o app .

# Set the environment variable for the server to run on
ENV PORT=8080

# Expose the port on which the server will run
EXPOSE 8080

# Command to run both HTTP and WebSocket servers
CMD ["bash", "-c", "go run main.go --server=http & go run main.go --server=websocket"]