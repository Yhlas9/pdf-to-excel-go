
BINARY_NAME=pdf-to-excel

MAIN=./cmd/main.go

.PHONY: all build run clean

build:
	go build -o $(BINARY_NAME) $(MAIN)

run: build
	./$(BINARY_NAME)

all: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)
