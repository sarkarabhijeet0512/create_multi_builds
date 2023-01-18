all: build

build:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 .
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 .
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 .
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64.exe .

clean:
	rm -f bin/*

