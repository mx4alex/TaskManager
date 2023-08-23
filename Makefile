.PHONY: run build
run: 
	go run ./cmd/app/main.go
build:
	go build -o ./build/taskmanager ./cmd/app/main.go
migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
test:
	go test -v ./...
swag:  
	swag init -g cmd/app/main.go
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           api/task/task.proto