BUILD_DEST_DIR ?= build

TARGET ?= $(shell cat VERSION | grep 'target=' | awk -F '=' '{print $$2}')
APP_VERSION ?= $(shell  cat VERSION | grep 'app=' | awk -F '=' '{print $$2}')
CHART_VERSION ?= $(shell cat VERSION | grep 'chart=' | awk -F '=' '{print $$2}')
yamls = $(wildcard kustomize/*.yaml)

.PHONY: build helm
build:
	mkdir -p ${BUILD_DEST_DIR}
	@echo "building ${BUILD_DEST_DIR}/${TARGET} ..."
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	CGO_ENABLED=0 go build -ldflags "-X main.VERSION=$(APP_VERSION) -X main.PROGRAM=$(TARGET)" \
		-o ${BUILD_DEST_DIR}/${TARGET} main.go

.PHONY: bindata
bindata:
	go generate bindata/plugins/sc-bindata.go

.PHONY: docker
docker:
	@echo "build docker image hub.expvent.com.cn:1111/expvent/${TARGET}:${APP_VERSION}"
	docker buildx build -f Dockerfile --build-arg TARGET=${TARGET} --build-arg APP_VERSION=${APP_VERSION} \
		-t hub.expvent.com.cn:1111/expvent/${TARGET}:${APP_VERSION} --platform=linux/amd64,linux/arm64 . --push