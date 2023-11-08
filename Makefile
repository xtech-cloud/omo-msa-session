APP_NAME := xm-msa-session
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. grpc/proto/sess.proto

.PHONY: build
build:
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run:
	./bin/${APP_NAME}

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call omo.msa.session SessionService.CheckAvailable '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.g93kIDX3blIWfAbREV7Aq87N97rhEyBEhTMH_cxKSn0"}'

.PHONY: tester
tester:
	go build -o ./bin/ ./tester

.PHONY: dist
dist:
	mkdir -p dist
	rm -f dist/${APP_NAME}-${BUILD_VERSION}.tar.gz
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build . -t omo.msa.session:latest

.PHONY: updev
updev: 
	scp -P 2209 dist/${APP_NAME}-${BUILD_VERSION}.tar.gz root@192.168.1.10:/root/
.PHONY: upload2
upload2:
	scp -P 9099 dist/${APP_NAME}-${BUILD_VERSION}.tar.gz root@47.93.209.105:/root/

.PHONY: upload
upload:
	scp -P 9099 dist/${APP_NAME}-${BUILD_VERSION}.tar.gz root@101.200.166.80:/root/