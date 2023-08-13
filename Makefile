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
<<<<<<< HEAD
	swag init -g cmd/app/main.go
=======
	swag init -g cmd/app/main.go
>>>>>>> f0b38b762063932898c2dc0b0be090d8dad20f13
