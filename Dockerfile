# STEP 1: Build executable binary
FROM golang:1.9 AS builder
# Copy project into image
COPY . $GOPATH/src/github.com/bradford-hamilton/go-graphql-api
# Set working directory to /go-graphql-api which contains main.go
WORKDIR $GOPATH/src/github.com/bradford-hamilton/go-graphql-api
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/api

# STEP 2: Build a small image start from scratch
FROM scratch
# Expose port 4000
EXPOSE 4000
# Copy our static executable
COPY --from=builder /go/bin/api .
# Command is run when container starts
ENTRYPOINT ["./api"]
