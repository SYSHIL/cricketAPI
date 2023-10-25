# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /service2
WORKDIR /service2

# Copy the service2 code into the container
COPY . .

# Build the Go application
RUN go build -o service2

# Expose port 8081
EXPOSE 8082

# Run the service1 executable
CMD ["./service2"]
