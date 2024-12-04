go-scaffolding: main.go
	go build

clean:
	-@rm go-scaffolding
	-@echo ''
	-@git clean -dfn && echo '`git clean -df` を実行すれば上記のファイル群が削除されます。'

.PHONY: run help test

run:
	go run main.go

help:
	go run main.go -h

test:
	go run main.go scaffold -v -r companion_ad
	-@echo ""
	go run main.go scaffold -v -r companion_ad --config .go-scaffolding.yaml
