GC=go build
BIN= bin
SRC = src
SERVEROBJ = Database.go Message.go User.go Util.go
SERVEROBJ := $(addprefix $(SRC)/, $(SERVEROBJ))
CLIENTOBJ = 
CLIENTOBJ := $(addprefix $(SRC)/, $(CLIENTOBJ))
SERVER = Server.go
CLIENT = Client.go

all: server

client: $(SRC)/$(CLIENT)
	@echo Compiling Go Test Client
	@$(GC) -o $(BIN)/$@ $(SRC)/$(CLIENT) $(CLIENTOBJ)

server: $(SRC)/$(SERVER)
	@echo Compiling Server
	@$(GC) -o $(BIN)/$@ $(SRC)/$(SERVER) $(SERVEROBJ)

.PHONY: clean

clean:
	@echo Removing compiled files
	@rm  $(BIN) *~  -rf
