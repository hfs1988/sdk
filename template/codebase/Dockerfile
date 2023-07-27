# Start from golang base image
FROM golang:1.20-alpine

# Setup folders
RUN mkdir /app

WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

# Expose port 3000 to the outside world
EXPOSE 3000

# Run the executable
CMD [ "/build" ]