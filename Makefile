.PHONY: build run

build:
	if not exist bin mkdir bin
	go build -o bin\main.exe .

run: build
	bin\main.exe