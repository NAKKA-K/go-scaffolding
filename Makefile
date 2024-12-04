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

test: RESOURCE=companion_ad
test:
	go run main.go scaffold -v -r $(RESOURCE) --config .go-scaffolding.yaml
