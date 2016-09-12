all: client_rule server_rule

client_rule: ./client/client.go
	@echo Compiling Client
	@go build -o goclient ./client/client.go

server_rule: ./server/server.go
	@echo Compiling Server
	@go build -o goserver ./server/server.go

clean:
	@echo Removing compiled files
	@rm goclient goserver *~
