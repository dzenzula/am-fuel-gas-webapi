
BUILD_NAME ?= am-fuel-gas-webapi
DEST_DIR := deploy/
DEPLOY_FILES := ${BUILD_NAME} configuration/${BUILD_NAME}.conf.yml service/${BUILD_NAME}.service

build:
	go build -ldflags "-s -w" -o ${BUILD_NAME} main.go

run:
	./${BUILD_NAME}

#build_and_run: build run

compile:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ${BUILD_NAME} main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ${BUILD_NAME} main.go
	GOARCH=amd64 GOOS=window go build -ldflags "-s -w" -o ${BUILD_NAME} main.go

install: ${DEPLOY_FILES}
	mkdir -p ${DEST_DIR}
	for f in ${DEPLOY_FILES}; do echo $$f;  cp -f $$f ${DEST_DIR}; done

clean:
	go clean
	rm -r ${DEST_DIR}

