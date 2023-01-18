all: build

build:
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 .
	md5 bin/main-linux-amd64 > main-linux-amd64.txt 
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 .
	md5 bin/main-linux-386 > main-linux-386.txt
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 .
	md5 bin/main-darwin-amd64 > main-darwin-amd64.txt
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64.exe .
	md5 bin/main-windows-amd64.exe > main-windows-amd64.txt
    
clean:
	rm -f bin/*

