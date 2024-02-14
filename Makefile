gen-code:
	protoc .\proto\chat.proto --go-grpc_out=.
	protoc .\proto\chat.proto --go_out=.
run-server:
	go run .\grpc_server.go

install:
	yarn install
	npm install grpc