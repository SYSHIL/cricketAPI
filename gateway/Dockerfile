# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory to /gateway
WORKDIR /gateway

# Copy the API gateway code into the container
COPY . .

# Build the Go application
RUN go build -o gateway

# Expose port 8080
EXPOSE 8080

# Run the API gateway executable
CMD ["./gateway"]
