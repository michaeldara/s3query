# @author Michael Dara

# parent image
FROM golang:latest

# workspace directory
WORKDIR /app

# copy `go.mod` and `go.sum`
ADD go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source code
COPY . .

# Build the Go app
RUN go build -o bin/s3query .

# create volume
VOLUME [ "/app/shared" ]

# set entrypoint
ENTRYPOINT [ "./bin/s3query" ]

