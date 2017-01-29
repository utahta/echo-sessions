.PHONY: fmt test

fmt:
	@goimports -w *.go 

test:
	@go test -v -race

