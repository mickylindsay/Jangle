@echo off
IF NOT EXIST bin
   mkdir bin
go build -o bin/server.exe Server.go User.go Message.go Database.go Util.go Parse.go Communication.go

