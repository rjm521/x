.PHONY: fmt test

fmt:
	@gofmt -l -w -s .
	@goimports -l -w .

test:
	go test -race -v ./...
