local:
	go build -o wxlogin.out *.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wxlogin_linux.out *.go
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o wxlogin_mac.out *.go
clean:
	rm -rvf *.out *.log *.jpeg
