dev:
	go run main.go

dev_compile:
	go build -o ./bin/server && ./bin/server

deploy:
	./bin/server