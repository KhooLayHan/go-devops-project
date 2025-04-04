APP_NAME=go-devops-app
VERSION=$(shell git describe --tags --always --dirty)
GIT_COMMIT=$(shell git rev-parse HEAD)
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: build test clean docker-build docker-push

build:
	go build -ldflags "-X github.com/KhooLayHan/go-devops-project/internal/version.Version=$(VERSION) \
		-X github.com/KhooLayHan/go-devops-project/internal/version.GitCommit=$(GIT_COMMIT) \
		-X github.com/KhooLayHan/go-devops-project/internal/version.BuildTime=$(BUILD_TIME)" \
		-o bin/$(APP_NAME) ./cmd

test:
	go test -v ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t $(APP_NAME):$(VERSION) .

docker-push:
	@echo "Implement your docker push command here"