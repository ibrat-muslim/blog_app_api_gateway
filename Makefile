swag-init:
	swag init -g api/api.go -o api/docs

start:
	go run cmd/main.go

.PHONY:	start