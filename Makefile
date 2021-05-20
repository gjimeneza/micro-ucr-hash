build:
	go build -o bin/micro-ucr-hash main.go

all: 
	build

clean: 
	rm -rf bin