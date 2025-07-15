.PHONY: build test clean package

BINARY_NAME=alfred-timestamp
WORKFLOW_NAME=alfred-timestamp.alfredworkflow

build:
	go build -o $(BINARY_NAME) main.go

test:
	go test -v ./...

clean:
	rm -f $(BINARY_NAME)
	rm -f $(WORKFLOW_NAME)

package: build
	zip -j $(WORKFLOW_NAME) $(BINARY_NAME) info.plist icon.png

install: package
	open $(WORKFLOW_NAME)

.DEFAULT_GOAL := build