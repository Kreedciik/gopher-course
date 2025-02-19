
server-start:
	go run server/main.go
client-start:
	go run client/main.go
tidy:
	go mod tidy
gen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/telemedicine.proto