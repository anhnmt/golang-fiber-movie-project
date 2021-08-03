# ./Dockerfile

FROM golang:1.16-alpine AS builder

RUN apk add --no-cache git

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image
# and build the API server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Move to working directory (/run).
WORKDIR /run

RUN mkdir -p /storage && mkdir -p /storage/banner && mkdir -p /storage/poster

# Copy binary and config files from /build
# to root folder of scratch container.
COPY --from=builder ["/build/main", "/build/config.yml", "/"]

ENV SERVER_PORT=8000 \
    SERVER_HOST=0.0.0.0

# Command to run when starting the container.
ENTRYPOINT ["./main"]

# Export necessary port.
EXPOSE 8000