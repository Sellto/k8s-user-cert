all:
	GOOS=linux GOARCH=amd64 go build -o usercert.linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o usercert.macos-amd64
	GOOS=linux GOARCH=arm go build -o usercert.arm
