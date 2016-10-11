CREATE TABLE users(userid INT UNSIGNED, username VARCHAR(20), displayname VARCHAR(20), imagepath VARCHAR(50), passwordhash VARCHAR(50), salt VARCHAR(8), UNIQUE KEY (username), PRIMARY KEY (userid));

CREATE TABLE servers( serverid INT UNSIGNED, servername VARCHAR(20), ownerid INT UNSIGNED, FOREIGN KEY (ownerid) REFERENCES users (userid), PRIMARY KEY (serverid));

CREATE TABLE rooms ( roomid INT UNSIGNED, roomname VARCHAR(30), roomdescription VARCHAR(1024), serverid INT UNSIGNED, PRIMARY KEY (roomid));

CREATE TABLE friends( userid INT UNSIGNED, friendid INT UNSIGNED, serverid INT UNSIGNED, FOREIGN KEY (userid) REFERENCES users (userid), FOREIGN KEY (friendid) REFERENCES users (userid), FOREIGN KEY (serverid) REFERENCES servers (serverid));

CREATE TABLE members ( userid INT UNSIGNED, serverid INT UNSIGNED, FOREIGN KEY (userid) REFERENCES users (userid), FOREIGN KEY (serverid) REFERENCES servers (serverid));

CREATE TABLE messages (userid INT UNSIGNED, time INT UNSIGNED, messageid INT UNSIGNED, messagetext BLOB, serverid INT UNSIGNED, roomid INT UNSIGNED,  PRIMARY KEY (messageid), FOREIGN KEY (userid) REFERENCES users (userid), FOREIGN KEY (serverid) REFERENCES servers (serverid), FOREIGN KEY (roomid) REFERENCES rooms (roomid));
