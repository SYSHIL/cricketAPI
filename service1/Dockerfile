# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /service1
WORKDIR /service1

# Copy the service1 code into the container
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application

# Build the Go application for Linux without disabling CGO
RUN GOOS=linux go build -o ./service1


RUN chmod +x ./service1

# Expose port 8081
EXPOSE 8081

# Run the service1 executable
CMD ["./service1"]
