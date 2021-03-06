FROM golang:1.16-alpine AS builder

WORKDIR /go/src/github.com/bethanyj28/battlesnek
COPY . /go/src/github.com/bethanyj28/battlesnek
RUN go mod vendor && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app github.com/bethanyj28/battlesnek/cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/bethanyj28/battlesnek/app .

EXPOSE 8080

CMD ["./app"]
