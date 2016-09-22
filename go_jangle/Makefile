GC=go build
BIN= bin
_TEST = $(BIN)/_test
SRC = src

all: server

client: $(SRC)/_test/Client.go
	@echo Compiling Go Test Client
	@$(GC) -o $(BIN)/_test/$@ $<

server: $(SRC)/Server.go
	@echo Compiling Server
	@$(GC) -o $(BIN)/$@ $(SRC)/*.go

.PHONY: clean

clean:
	@echo Removing compiled files
	@rm  $(BIN) *~ -rf

