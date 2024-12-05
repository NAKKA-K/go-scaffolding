BIN := $(CURDIR)/tools/_bin

$(BIN)/%: go.mod go.sum tools/go.mod tools/go.sum
	cd tools && cat tools.go | awk -F'"' '/_/ {print $$2}' | grep $* | GOBIN=$(BIN) xargs -tI {} go install {}

lint: $(BIN)/golangci-lint .golangci.yaml
	$< run

fix: $(BIN)/golangci-lint .golangci.yaml
	$< run --fix

go-scaffolding: main.go
	go build

clean:
	-@rm go-scaffolding
	-@echo ''
	-@git clean -dfn && echo '`git clean -df` を実行すれば上記のファイル群が削除されます。'

.PHONY: run help test

RESOURCE :=

run:
	go run main.go -r $(RESOURCE)

help:
	go run main.go -h

test: RESOURCE=resource_snake_case
test:
	go run main.go scaffold -v -r $(RESOURCE) --config .go-scaffolding.yaml

