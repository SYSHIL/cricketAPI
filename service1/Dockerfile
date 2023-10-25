# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /service1
WORKDIR /service1

# Copy the service1 code into the container
COPY . .

# Build the Go application
RUN go build -o service1

# Expose port 8081
EXPOSE 8081

# Run the service1 executable
CMD ["./service1"]
