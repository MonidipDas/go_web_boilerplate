.PHONY: build run

build:
	if not exist bin mkdir bin
	go build -o bin\main.exe .\cmd\api\main.go

run: build
	bin\main.exe