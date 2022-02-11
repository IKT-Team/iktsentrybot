FILES = $(wildcard **.go)

build:
	GOARCH=amd64 GOOS=linux go build -o main

run:
	go run $(FILES)

clean:
	rm -f deploy.zip main debug

deploy: clean build
	zip deploy.zip main
