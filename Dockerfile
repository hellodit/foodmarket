############################
# STEP 1 build executable binary
############################
FROM golang:1.14 AS builder 

# Install git.
# Git is required for fetching the dependencies.
RUN apt update && apt install git -y && apt autoclean -y && apt autoremove -y 
# Create appuser.
ENV USER=appuser
ENV UID=10001 
# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /foodmarket
# Copy go mod and sum files 
COPY go.mod go.sum ./

RUN go mod download 

COPY . . 


# Fetch dependencies.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /main.go
