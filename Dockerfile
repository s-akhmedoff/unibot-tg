FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /unibot/src

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o unibot cmd/main.go

FROM scratch

COPY --from=builder /unibot/src/unibot /unibot/bin/unibot

ENTRYPOINT ["/unibot/bin/unibot"]