
server-start:
	go run server/server.go
client-start:
	go run client/client.go
tidy:
	go mod tidy
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/heart-rate.proto