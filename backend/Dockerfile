FROM golang:latest

RUN apt-get update
RUN apt-get install lsof

WORKDIR /app

COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

EXPOSE 8080