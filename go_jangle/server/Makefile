GC = go build
BIN = bin
SRC = src
SERVEROBJ = Database.go Message.go User.go Util.go Communication.go Parse.go Crypt.go
SERVEROBJ := $(addprefix $(SRC)/, $(SERVEROBJ))
SERVER = Server.go

all: server

server: $(SRC)/$(SERVER)
	@echo Compiling Server
	@$(GC) -o $(BIN)/$@ $(SRC)/$(SERVER) $(SERVEROBJ)

.PHONY: clean

clean:
	@echo Removing compiled files
	@rm  $(BIN) *~  -rf
