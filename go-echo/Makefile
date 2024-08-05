# Makefile
TARGET = "everai-test-go"
APP_VERSION = "v0.0.2"

.PHONY: docker
docker:
	@echo "build docker image hub.expvent.com.cn:1111/expvent/${TARGET}:${TARGET}"
	docker buildx build -f Dockerfile --build-arg TARGET=${TARGET} --build-arg APP_VERSION=${APP_VERSION} \
		-t hub.expvent.com.cn:1111/expvent/${TARGET}:${APP_VERSION} --platform=linux/amd64,linux/arm64 . --push

all: docker