# Building the binary of the App
FROM golang:1.16 AS build

# `base-project` should be replaced with your project name
WORKDIR /go/src/base-project

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .


# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest

WORKDIR /app

# Create the `public` dir and copy all the assets into it
RUN mkdir ./static
COPY ./static ./static

# `base-project` should be replaced here as well
COPY --from=build /go/src/base-project/app .

# Exposes port 3000 because our program listens on that port
EXPOSE 3000

CMD ["./app"]