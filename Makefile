.PHONY: coverage test

test:
	@go test -coverprofile=coverage.out ./...

coverage:
	@go tool cover -html=coverage.out
