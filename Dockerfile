# Start from golang base
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Asdita Prasetya <asditap@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
# RUN apk update && apk add --no-cache git

RUN mkdir /foodmarket

# Set the current working directory inside the container
# Copy go mod and sum files
COPY . /foodmarket

WORKDIR /foodmarket
RUN go mod tidy
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download
# Build the Go app

RUN go build -o main


# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /foodmarket/main .
COPY --from=builder /foodmarket/config.yaml .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]