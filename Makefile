all: run

run: build
	./what GNU

build: main.go
	go build

clean:
	rm -f what
