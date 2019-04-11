all: build

test: build
	./acro GNU

build: main.go
	go build

clean:
	rm -f acro
