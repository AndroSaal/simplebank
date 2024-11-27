SRC_TO_PROTO=proto
SRC_TO_GATEWAY_AUTH_PROTO=$(SRC_TO_PROTO)/gateway-auth/v1/gateway-auth.proto
SRC_TO_AUTH_SERVICE=./services/auth/pb

prAuth: install_packages
	protoc -I $(SRC_TO_PROTO) $(SRC_TO_GATEWAY_AUTH_PROTO) \
	--go_out=$(SRC_TO_AUTH_SERVICE) --go_opt=paths=source_relative \
	--go-grpc_out=$(SRC_TO_AUTH_SERVICE) --go-grpc_opt=paths=source_relative \
	--plugin=$$(which protoc-gen-go) --plugin=$$(which protoc-gen-go-grpc)

install_packages:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
