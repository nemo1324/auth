.SILENT:

# Конфигурация
GRPC_PORT=44044 \
POSTGRES_HOST=postgres \
POSTGRES_PORT=5432 \
POSTGRES_USER=postgres \
POSTGRES_DB_NAME=postgres \
POSTGRES_PASSWORD=qwerty \
TOKEN_TTL=1h

# Путь к директории проекта
PROJECT_DIR := $(PWD)

run:
	$(CONFIG_VARS) \
	go run cmd/auth/main.go


protoc:
	protoc -I protos/proto protos/proto/auth/auth.proto --go_out=protos/gen/go/ --go_opt=paths=source_relative --go-grpc_out=protos/gen/go/ --go-grpc_opt=paths=source_relative
