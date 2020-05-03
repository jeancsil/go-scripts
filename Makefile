build:
	go build -o bin/head commands/head.go

run:
	go run main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/head-freebsd-386 commands/head.go
	GOOS=linux GOARCH=386 go build -o bin/head-linux-386 commands/head.go
	GOOS=windows GOARCH=386 go build -o bin/head.exe commands/head.go
