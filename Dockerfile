############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/ctfrancia/yourticket
COPY . .
# Fetch dependencies.
# Using go get.
RUN go mod download
# Build the binary.
RUN go build -o /go/bin/yourticket ./cmd/api
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/yourticket /go/bin/yourticket
# Run the hello binary.
ENTRYPOINT ["/go/bin/yourticket"]
