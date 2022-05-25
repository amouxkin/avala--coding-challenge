test:
	go test ./...

dep:
	go mod download

vet:
	go vet

clean:
	go clean

lint:
	 golangci-lint run
