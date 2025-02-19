generate:
	go generate ./...

lint:
	golint -set_exit_status ./...

test: generate lint
	go test -race -cover ./...
	go run ./cmd/tengo -resolve ./testdata/cli/test.tengo

fmt:
	go fmt ./...

cli:
	go build -o tengo cmd/tengo/main.go
