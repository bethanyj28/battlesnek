FROM golang:1.16-alpine AS builder

WORKDIR /go/src/github.com/bethanyj28/battlesnek
COPY . /go/src/github.com/bethanyj28/battlesnek
RUN go mod vendor && go mod tidy
ENV PORT=$PORT
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./...

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/bethanyj28/battlesnek/app .

EXPOSE $PORT

CMD ["./app"]
