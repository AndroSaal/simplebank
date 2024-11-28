PATH_TO_PROTO=proto
PATH_TO_GATEWAY_AUTH_PROTO=$(PATH_TO_PROTO)/gateway-auth/v1/gateway-auth.proto
PATH_TO_AUTH_SERVICE=./services/auth/pb
PATH_TO_AUTH_SERVICE_FOLDER=./services/auth


prAuth: install_packages
	protoc -I $(PATH_TO_PROTO) $(PATH_TO_GATEWAY_AUTH_PROTO) \
	--go_out=$(PATH_TO_AUTH_SERVICE) --go_opt=paths=source_relative \
	--go-grpc_out=$(PATH_TO_AUTH_SERVICE) --go-grpc_opt=paths=source_relative \
	--plugin=$$(which protoc-gen-go) --plugin=$$(which protoc-gen-go-grpc)

install_packages:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

all : prAuth
	make -C $(PATH_TO_AUTH_SERVICE_FOLDER)