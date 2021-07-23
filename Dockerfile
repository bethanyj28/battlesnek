FROM golang:1.16-alpine AS builder

WORKDIR /go/src/github.com/bethanyj28/go-api-starter
COPY . /go/src/github.com/bethanyj28/go-api-starter
RUN go mod vendor && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./...

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/bethanyj28/go-api-starter/app .

ENTRYPOINT ["./app"]

EXPOSE 8080
