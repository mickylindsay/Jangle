GC=go build
BIN= bin
_TEST = $(BIN)/_test
SRC = src

all: server

client: $(SRC)/_test/client.go
	@echo Compiling Go Test Client
	@$(GC) -o $(BIN)/_test/$@ $<

server: $(SRC)/server.go
	@echo Compiling Server
	@$(GC) -o $(BIN)/$@ $(SRC)/*.go

.PHONY: clean

clean:
	@echo Removing compiled files
	@rm  $(BIN) *~ -rf

