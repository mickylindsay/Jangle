echo off

go build -o bin/server.exe src/Server.go src/User.go src/Message.go src/Database.go src/Util.go src/Parse.go src/Communication.go src/Crypt.go
