build:
	go build -o bin/head commands/head.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/head commands/head.go
	GOOS=linux GOARCH=amd64 go build -o bin/head-linux-386 commands/head.go
	GOOS=windows GOARCH=amd64 go build -o bin/head.exe commands/head.go
