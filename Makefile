BINARY_NAME=app
INSTALL_PATH=/usr/bin/

${BINARY_NAME}:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} cmd/main.go

install: ${BINARY_NAME}
	@echo; sudo install ${BINARY_NAME} -v -s -m 0770 -o root -g root -t ${INSTALL_PATH}

clean:
	@echo; go clean
	@echo; rm -fv ${BINARY_NAME}