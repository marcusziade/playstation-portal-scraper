# Use an official Go runtime as a parent image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Fetch dependencies
RUN go mod init playstation-scraper
RUN go get -d -v ./...

# Build the Go app
RUN go build -o main .

# Run the application when the container starts
CMD ["./main"]
