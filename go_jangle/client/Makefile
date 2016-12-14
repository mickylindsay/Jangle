GC = go build
BIN = bin
SRC = src
CLIENTOBJ =
CLIENTOBJ := $(addprefix $(SRC)/, $(CLIENTOBJ))
CLIENT = Client.go

all: client

client: $(SRC)/$(CLIENT)
	@echo Compiling Go Test Client
	@$(GC) -o $(BIN)/$@ $(SRC)/$(CLIENT) $(CLIENTOBJ)

.PHONY: clean

clean:
	@echo Removing compiled files
	@rm  $(BIN) *~  -rf
