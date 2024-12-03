go-scaffolding: main.go
	go build

clean:
	-@rm go-scaffolding

.PHONY: run help

run:
	go run main.go

help:
	go run main.go -h
