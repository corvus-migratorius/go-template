BINARY_NAME=app
INSTALL_PATH=/usr/bin/
TAG=$(shell git describe --abbrev=0 || echo "0.0.1")
HASH=$(shell git rev-parse --verify --short HEAD)

${BINARY_NAME}:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-X main.Version=${TAG}-${HASH}" -o ${BINARY_NAME} cmd/main.go

install: ${BINARY_NAME}
	@echo; sudo install ${BINARY_NAME} -v -s -m 0770 -o root -g root -t ${INSTALL_PATH}

clean:
	@echo; go clean
	@echo; rm -fv ${BINARY_NAME}