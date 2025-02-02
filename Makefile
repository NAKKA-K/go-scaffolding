BIN := $(CURDIR)/tools/_bin

.PHONY: lint fix test clean run help

$(BIN)/%: go.mod go.sum tools/go.mod tools/go.sum
	cd tools && cat tools.go | awk -F'"' '/_/ {print $$2}' | grep $* | GOBIN=$(BIN) xargs -tI {} go install {}

lint: $(BIN)/golangci-lint .golangci.yaml
	$< run

fix: $(BIN)/golangci-lint .golangci.yaml
	$< run --fix

test:
	go test -v ./...

go-scaffolding: main.go
	go build

clean:
	-@rm go-scaffolding
	-@echo ''
	-@git clean -dfn && echo '`git clean -df` を実行すれば上記のファイル群が削除されます。'

RESOURCE :=

run:
	go run main.go scaffold -r $(RESOURCE)

help:
	go run main.go -h

o:
	go run main.go scaffold api_test -r test_resource

