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

FROM scratch

COPY --from=builder /go/bin/yourticket /go/bin/yourticket

ENTRYPOINT ["/go/bin/yourticket"]
EXPOSE 8080
