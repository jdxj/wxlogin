mac:
	go build -o main.out *.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main_linux.out *.go
clean:
	rm -v *.out *.log
