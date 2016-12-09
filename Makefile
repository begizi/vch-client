.PHONY: build

build:
	@cd cmd/vch && GOARCH=arm GOOS=linux CGO_ENABLED=0 go build -o ../../bin/vch.pi
	@cd cmd/vch && GOARCH=amd64 CGOOS=linux CGO_ENABLED=0 go build -o ../../bin/vch.linux
	@cd cmd/vch && CGO_ENABLED=0 go build -o ../../bin/vch
