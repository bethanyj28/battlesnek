build:
	docker build -t bethanyj28/battlesnek .
run:
	docker run --rm -p 8080:8080 bethanyj28/battlesnek
test:
	go test ./...
vendor:
	go mod vendor && go mod tidy
env:
	touch environment.env
