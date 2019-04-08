all: run

run: build
	./xmltest

build: main.go
	go build

clean:
	rm -f xmltest
