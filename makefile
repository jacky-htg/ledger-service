gen:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/*/*.proto
	
migrate:
	go run cmd/cli.go migrate
	
seed:
	go run cmd/cli.go seed
	
server:
	go run server.go
	
.PHONY: gen migrate seed server