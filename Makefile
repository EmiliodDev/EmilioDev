build:
	@templ generate && go build -o bin/emiliodev cmd/emiliodev/main.go

test:
	@go test -v ./...

run: build
	@./bin/emiliodev
