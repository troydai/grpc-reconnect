.PHONY: up
up:
	@ docker compose up --build

.PHONY: gen
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	protos/echo.proto

.PHONY: tools
tools:
	@ brew install protobuf
	@ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: build
build:
	@ mkdir -p artifacts/grpc-reconnect
	@ go build -o artifacts/grpc-reconnect/server cmd/server/main.go
	@ go build -o artifacts/grpc-reconnect/client cmd/client/main.go
