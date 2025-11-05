.PHONY: gogen

gogen:
	go generate ./...

.PHONY: test

test:
	go test ./...

.PHONY: vet

vet:
	go vet ./...
