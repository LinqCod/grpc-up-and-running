proto_client:
	protoc -I proto ./proto/*.proto --go_out=./productInfo/client/ecommerce --go_opt=paths=source_relative --go-grpc_out=./productInfo/client/ecommerce --go-grpc_opt=paths=source_relative

proto_server:
	protoc -I proto ./proto/*.proto --go_out=./productInfo/server/ecommerce --go_opt=paths=source_relative --go-grpc_out=./productInfo/server/ecommerce --go-grpc_opt=paths=source_relative

.PHONY: proto_server, proto_client