build:
	docker build -t bethanyj28/go-api-starter .
run:
	docker run --rm -p 8080:8080 bethanyj28/go-api-starter
test:
	go test ./...
vendor:
	go mod vendor && go mod tidy
env:
	touch environment.env
