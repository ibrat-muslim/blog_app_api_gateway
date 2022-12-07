CURRENT_DIR=$(shell pwd)

swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

proto-gen:
	rm -rf genproto
	./scripts/gen-proto.sh ${CURRENT_DIR}

.PHONY:	start