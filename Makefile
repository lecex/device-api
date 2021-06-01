.PHONY: git
git:
	git add .
	git commit -m"自动提交 git 代码"
	git push
.PHONY: tag
tag:
	git push --tags
.PHONY: micro
micro:
	micro api --enable_rpc=true

.PHONY: proto
proto:
	protoc -I . --micro_out=. --gogofaster_out=. proto/device/device.proto
	
.PHONY: docker
docker:
	docker build -f Dockerfile  -t device-api .
.PHONY: run
run:
	go run main.go
test:
	go test main_test.go -test.v