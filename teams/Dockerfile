# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /service1
WORKDIR /teams

# Copy the service1 code into the container
COPY . .

# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application

RUN CGO_ENABLED=0 GOOS=linux go build -o ./teams

RUN chmod +x ./teams
# Expose a port (if your Go application listens on a port)
EXPOSE 8083

# Run the Go application
CMD ["./teams"]

EXPOSE 8083

# Run the service1 executable
CMD ["./teams"]
