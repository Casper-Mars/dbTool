VERSION = 0_4_1

buildLinux:
	go build -o DBTool-$(VERSION)-linux_amd64 -x -v ./

buildMac:
	go build -o DBTool-$(VERSION)-darwin_amd64 -x -v ./

