default:
	@go build -o "./build/game-linux"
	@./build/game-linux

run:
	@go run main.go
