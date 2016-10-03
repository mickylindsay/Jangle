CREATE TABLE users(userid INT, username VARCHAR(20), displayname VARCHAR(20), imagepath VARCHAR(50), passwordhash VARCHAR(50), salt VARCHAR(8), UNIQUE KEY (username), PRIMARY KEY (userid));

CREATE TABLE servers( serverid INT, servername VARCHAR(20), ownerid INT, FOREIGN KEY (ownerid) REFERENCES users (userid), PRIMARY KEY (serverid));

CREATE TABLE rooms ( roomid INT, roomname VARCHAR(30), roomdescription VARCHAR(1024), serverid INT, PRIMARY KEY (roomid));

CREATE TABLE friends( userid INT, friendid INT,	serverid INT, FOREIGN KEY (userid) REFERENCES users (userid), FOREIGN KEY (friendid) REFERENCES users (userid),	FOREIGN KEY (serverid) REFERENCES servers (serverid));

CREATE TABLE members ( userid INT, serverid INT, FOREIGN KEY (userid) REFERENCES users (userid), FOREIGN KEY (serverid) REFERENCES servers (serverid));

CREATE TABLE messages (	userid INT, time INT UNSIGNED, messageid INT, messagetext VARCHAR(1024), serverid INT, PRIMARY KEY (messageid),	FOREIGN KEY (userid) REFERENCES users (userid), FOREIGN KEY (serverid) REFERENCES servers (serverid));
