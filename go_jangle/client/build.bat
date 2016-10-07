@echo off
IF NOT EXIST bin
   mkdir bin
set obj = User.go Message.go Database.go Util.go
IF "%1"=="-c" 
   go build -o bin/server.exe Server.go %obj%
ELSE
   go build -o bin/client.exe Client.go %obj%
