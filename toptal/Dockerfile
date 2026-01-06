# First of all, we are doing multi-stage build

# Go tools + alpine for builder stage
# Name it AS builder so that you can use it later in the file
FROM golang:1.22-alpine AS builder

# --- BUILDER STAGE ---

# Set the working directory inside the container
WORKDIR /app

# Copy only module files from host into a current working directory of a container "./"
COPY go.mod go.sum ./

# Download the Go module dependencies. Done at build stage
RUN go mod download

# Copy the source code into the container
# First copy mod files, install so that if code changes no re-inatall is needed
COPY . .

# disable CGO - static binary that runes fine on Alpine
# GOOS - linux target build
# -a - force rebuild all packages
# -installsuffix cgo - keep a separate build cache suffix for cgo settings 
# -o app - set output name to "app"
# ./cmd/main.go - location of the file we are building
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/main.go

# --- RUNTIME STAGE ---

# This one prepares runtime

# Alpine for runtime stage
FROM alpine:latest

# Runs once per image build
RUN apk update && apk upgrade && apk add bash

# Reduce image size: delete cacked stuff and temp files to reduce image size
# var/cache/apk - place where apache caches downloads
# /tmp - standard temporary directory
RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# Avoid running code as a root user
# Create a new user, switch to it
RUN adduser -D appuser
USER appuser

# Set (Again) the working directory inside the container
# because we are in a new stage - this should be done once again
WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/cmd/wait-for-it.sh .
COPY --from=builder /app/internal/app/migrations ./migrations

# Expose the port that the application listens on
EXPOSE 8080

# Run the binary when the container starts
CMD ["./app"]