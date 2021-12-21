# Start from golang base image
FROM golang:1.17 as builder

# Set the current working directory inside the container
WORKDIR /build

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

# Build the Go app
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server main.go

# Start a new stage from busybox
FROM alpine:3.11

WORKDIR /dist
RUN apk update && apk add ca-certificates

# Copy the build artifacts from the previous stage
COPY --from=builder /build/server .

# Run the executable
CMD ["./server"]