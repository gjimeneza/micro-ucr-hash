build:
	go build -o bin/micro-ucr-hash main.go

all: 
	build

test:
	go test -v ./...

clean: 
	rm -rf bin