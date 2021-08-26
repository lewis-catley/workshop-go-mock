run-basic:
	go run basic/main.go

run-better:
	go run better/*.go

create-mocks:
	go generate ./...

lint:
	golint ./...

test: create-mocks
	go test -cover ./better/...

test-html: create-mocks
	go test -coverprofile=coverage.out ./better/...
	go tool cover -func=coverage.out
