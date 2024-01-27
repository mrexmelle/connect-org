
PROJECT_NAME=connect-org
VERSION=0.2.0
IMAGE_NAME=ghcr.io/mrexmelle/$(PROJECT_NAME)
GO_SOURCES=$(shell find . -name '*.go' -not -path "./vendor/*")

$(PROJECT_NAME): $(GO_SOURCES)
	go build -o $@ ./cmd/main.go

clean:
	rm -rf $(PROJECT_NAME)

distclean:
	rm -rf $(PROJECT_NAME) docs

docker-build:
	docker build -t $(IMAGE_NAME):$(VERSION) .

docker-push:
	docker push $(IMAGE_NAME):$(VERSION)

docs:
	swag init --parseDependency -g cmd/main.go

test:
	go test ./internal/...

all: $(PROJECT_NAME)