package com.jangle.communicate;

import java.io.IOException;
import java.net.UnknownHostException;
import java.util.Arrays;

import com.jangle.client.*;
import com.jangle.communicate.CommUtil.*;

public class Client_ParseData implements IPARSER {

	private Client mClient;
	private Client_Communicator Comm;

	// Variables used when recieving data from the server. These are used as
	// temporary storage, and are not guaranted to hold any value
	private LoginResult loginResult;
	private int UserID;
	private String DisplayName;
	private int numMessagesRecieved;
	private String ServerDisplayName;
	private String RoomDisplayName;
	private String IP;

	/**
	 * Create a parser object with no Client_Commmunicator attached to it.
	 * 
	 * @param Clie
	 *            The client object this communicator references
	 */
	public Client_ParseData(Client Clie) {
		this.mClient = Clie;

	}

	/**
	 * Create a parser object with a Client_Communicator attached to it.
	 * 
	 * @param Clie
	 *            The client object this communicator references
	 * @param Host
	 *            The IP address of the server
	 * @param port
	 *            port to communicate though with the server
	 * @throws UnknownHostException
	 * @throws IOException
	 */
	public Client_ParseData(Client Clie, String Host, int port) throws UnknownHostException, IOException {
		this.mClient = Clie;
		this.Comm = new Client_Communicator(this, Host, port);
	}

	/**
	 * Send a Message Objects's info to the server, as per the message
	 * specification
	 */
	public void sendMessage(Message mess) throws IOException {
		Comm.sendToServer(mess.getByteArray());
	}

	/**
	 * Figure out what the data that was received is.
	 * 
	 * @param data
	 *            the character array to parse, and figure out what it is
	 */
	public void parseData(byte[] data) {

		if (data[0] == CommUtil.MESSAGE_FROM_SERVER) {
			Message newMess = new Message(data);
			//System.out.println("Server id: " + newMess.getServerID() + " channelid: " + newMess.getChannelID());
			if (!mClient.isDuplicateMessage(newMess))
				mClient.addMessage(newMess, newMess.getServerID(), newMess.getChannelID());
			/*
			 * //This code adds user to ui if not added already for(int i = 0; i
			 * < mClient.getUsers().size(); i++) { if (newMess.getUserID() ==
			 * mClient.getUsers().get(i).getId()){ return; } }
			 * 
			 * User newUser = new User("" + newMess.getUserID(),
			 * newMess.getUserID()); mClient.getUsers().add(newUser); try {
			 * requestDisplayName(newUser); } catch (IOException e) {
			 * e.printStackTrace(); }
			 */

		}

		else if (data[0] == CommUtil.LOGIN_SUCCESS) {
			mClient.setLoginResult(LoginResult.SUCESS);
			mClient.setUserID(CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length)));
			return;
		}

		else if (data[0] == CommUtil.LOGIN_FAIL) {
			mClient.setLoginResult(LoginResult.FAIL);
			return;
		}

		else if (data[0] == CommUtil.CREATE_USER_FAIL) {
			mClient.setLoginResult(LoginResult.NAME_TAKEN);
			return;
		}

		else if (data[0] == CommUtil.RECIEVE_DISPLAY_NAME) {
			// TODO: Changing
			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
			String newDiplay = new String(Arrays.copyOfRange(data, 5, data.length));
			for (int i = 0; i < mClient.getUsers().size(); i++) {
				if (id == mClient.getUsers().get(i).getId()) {
					mClient.getUsers().get(i).setDisplayName(newDiplay);
					return;
				}
			}
			// If user is not already added we add them
			mClient.getUsers().add(new User(newDiplay, id));
		}

		else if (data[0] == CommUtil.RECIEVE_USERID) {

			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
			User tmp = new User("" + id, id);

			for (int i = 0; i < mClient.getUsers().size(); i++) {
				if (id == mClient.getUsers().get(i).getId())
					return;
			}

			mClient.addUser(tmp);

			try {
				requestDisplayName(tmp);
			} catch (IOException e) {
				e.printStackTrace();
			}

		}

		else if (data[0] == CommUtil.RECIEVE_SERVER_ID) {
			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
			Server newServer = new Server(id);
			mClient.addServer(newServer);
			requestServerDisplayName(newServer);
			requestAllRoomsInServer(newServer);
            requestServerIcon(newServer);
		}

		else if (data[0] == CommUtil.RECIEVE_SERVER_DISPLAY_NAME) {
			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
			String displayName = new String(Arrays.copyOfRange(data, 5, data.length));

			if (mClient.getServer(id) != null)
				mClient.getServer(id).setName(displayName);

		}

		else if (data[0] == CommUtil.RECIEVE_ROOM_ID) {
			int sId = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
			int chId = CommUtil.byteToInt(Arrays.copyOfRange(data, 5, 9));

			Channel newChannel = new Channel(chId);
			mClient.getServer(sId).addChannel(newChannel);
            if (sId == mClient.getCurrentServerID() && mClient.findUser(chId + 1000) == null){
                mClient.addUser(new User(newChannel));
            }


			try {
				requestRoomDisplayName(sId, chId);
			} catch (IOException e) {
				e.printStackTrace();
			}

		}

        else if (data[0] == CommUtil.RECIEVE_USER_LOCATION) {
            //System.out.println("Recieved user location change");
            int serverID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
            int channelID = CommUtil.byteToInt(Arrays.copyOfRange(data, 5, 9));
            int userID = CommUtil.byteToInt(Arrays.copyOfRange(data, 9, 13));
            mClient.updateUserPosition(userID, serverID, channelID);
            mClient.setLocationChanged(true);
        }

		else if (data[0] == CommUtil.RECIEVE_ROOM_DISPLAY_NAME) {
			int sId = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
			int chId = CommUtil.byteToInt(Arrays.copyOfRange(data, 5, 9));
			String displayName = new String(Arrays.copyOfRange(data, 9, data.length));

			mClient.getServer(sId).getChannel(chId).setName(displayName);
		}
		else if (data[0] == CommUtil.RECIEVE_USER_IP) {
			byte[] address = new byte[data.length - 4];
			for (int i = 0; i < address.length; i++) {
				address[i] = data[i + 4];
			}
			IP = new String(address);

			int loc = IP.indexOf(':');
			IP = IP.substring(0, loc);
		}

		else if (data[0] == CommUtil.RECIEVE_USER_STATUS) {

			int userID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));

			User user = mClient.findUser(userID);
			if (user == null) {
                User newUser = new User("" + userID, userID);
				mClient.addUser(newUser);
                try {
                    requestDisplayName(newUser);
                } catch (IOException e) {
                    e.printStackTrace();
                }
                return;
			}

			if (data[5] == (byte) 0) {
				user.setStatus(CommUtil.UserStatus.OFFLINE);
			}
			else if (data[5] == (byte) 1) {
				user.setStatus(CommUtil.UserStatus.ONLINE);
			}
			else {
				user.setStatus(CommUtil.UserStatus.AWAY);
			}

			if (data[6] == (byte) 0) {
				user.setIsMuted(false);
			}
			else {
				user.setIsMuted(true);
			}

			if (data[7] == (byte) 0) {
				user.setVoiceStatus(false);
			}
			else {
				user.setVoiceStatus(true);
			}
		}

        else if(data[0] == CommUtil.RECIEVE_SERVER_ICON) {
            int serverID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
            String Img = new String(Arrays.copyOfRange(data, 5, data.length));

            mClient.getServer(serverID).setAvatarURL(Img);
        }

        else if(data[0] == CommUtil.RECIEVE_USER_ICON) {
            int userID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, 5));
            String img = new String(Arrays.copyOfRange(data, 5, data.length));

            mClient.findUser(userID).setAvatar(img);
        }

	}

    private void requestServerIcon(Server newServer) {
        byte[] toSend = new byte[5];
        toSend[0] = CommUtil.REQUEST_SERVER_ICON;

        byte[] idInByte = CommUtil.intToByteArr(newServer.getId());

        for (int i = 0; i < idInByte.length; i++) {
            toSend[i + 1] = idInByte[i];
        }

        try {
            Comm.sendToServer(toSend);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
	 * Submits a login request to the server. If the login is a success, the
	 * user ID of the client that it passed to this parser when initalized will
	 * get set to the user's userID
	 * 
	 * @param Username
	 *            The username for the user
	 * @param Password
	 *            The password for the user
	 * @return If the Login was a success
	 * @throws IOException
	 */
	public void submitLogIn(String Username, String Password) throws IOException {

		byte[] data = new byte[20 + Password.length() + 1];
		mClient.setLoginResult(LoginResult.TIMEOUT);
		mClient.setLoginTime(System.currentTimeMillis());
		int place = 0;

		data[0] = CommUtil.LOGIN;
		place++;

		for (int i = 0; i < Username.length(); i++) {
			data[place] = Username.getBytes()[i];
			place++;
		}

		for (; place < 21; place++) {
			data[place] = (byte) 0;
		}
		for (int i = 0; i < Password.length(); i++) {
			data[place] = Password.getBytes()[i];
			place++;
		}

		Comm.sendToServer(data);

		return;
	}

	/**
	 * Submits a create user request. If the creation request is a success, the
	 * userID of the client that is given at this parser instantiation will be
	 * set to the userID given by the server
	 * 
	 * @param Username
	 * @param Password
	 * @return
	 * @throws IOException
	 */
	public void createUserInServer(String Username, String Password) throws IOException {

		byte[] data = new byte[20 + Password.length() + 1];
		mClient.setLoginResult(LoginResult.TIMEOUT);
		mClient.setLoginTime(System.currentTimeMillis());
		int place = 0;

		data[0] = CommUtil.CREATE_USER;
		place++;

		for (int i = 0; i < Username.length(); i++) {
			data[place] = Username.getBytes()[i];
			place++;
		}

		for (; place < 21; place++) {
			data[place] = (byte) 0;
		}
		for (int i = 0; i < Password.length(); i++) {
			data[place] = Password.getBytes()[i];
			place++;
		}

		Comm.sendToServer(data);

		return;
	}

	/**
	 * Request block of 50 messages from the server
	 * 
	 * @param offSet
	 *            Which block of 50 to
	 * @return
	 * @throws IOException
	 */
	public void request50MessagesWithOffset(int offSet) throws IOException {

		numMessagesRecieved = 0;

		byte[] test = new byte[2];
		test[0] = CommUtil.REQUEST_N_MESSAGES;
		test[1] = (byte) offSet;

		Comm.sendToServer(test);
	}

	/**
	 * Sends request of all of the servers the given user is a member with (33)
	 * 
	 * @throws IOException
	 */
	public void requestAllServers(User user) throws IOException {

		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_ALL_SERVERID;

		byte[] tmp = CommUtil.intToByteArr(user.getId());

		for (int i = 0; i < tmp.length; i++) {
			toServer[i + 1] = tmp[i];
		}

		Comm.sendToServer(toServer);

	}

	public void requestDisplayName(User user) throws IOException {

		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_DISPLAY_NAME;

		byte[] idInByte = CommUtil.intToByteArr(user.getId());

		for (int i = 0; i < idInByte.length; i++) {
			toServer[i + 1] = idInByte[i];
		}

		Comm.sendToServer(toServer);

	}

	/**
	 * Set a new display name for the logged in user
	 * 
	 * @param user
	 *            The name to set the new user as
	 * @throws IOException
	 *             If cannot send the data to the server.
	 */
	public void setNewDisplayNameOnServer(String user) throws IOException {
		byte[] toServer = new byte[user.length() + 1];
		byte[] nameAsByte = user.getBytes();
		toServer[0] = CommUtil.SEND_NEW_DISPLAY_NAME;

		for (int i = 0; i < nameAsByte.length; i++) {
			toServer[i + 1] = nameAsByte[i];
		}

		Comm.sendToServer(toServer);
	}

	/**
	 * Request all of the userID that are members of the connected server (35)
	 * 
	 * @throws IOException
	 */
	public void requestAllUsersTiedToServer() throws IOException {
		byte[] toServer = new byte[1];
		toServer[0] = CommUtil.REQUEST_ALL_USERID;
		Comm.sendToServer(toServer);
	}

	// TODO need to test this
	/**
	 * get a list of all of the room IDs in the room
	 * 
	 * @param serverID
	 * @throws IOException
	 */
	public void getRoomIDInServer(int serverID) throws IOException {
		byte[] toServer = new byte[5];
		byte[] nameAsByte = CommUtil.intToByteArr(serverID);
		toServer[0] = CommUtil.REQUEST_ALL_ROOM_ID;

		for (int i = 0; i < nameAsByte.length; i++) {
			toServer[i + 1] = nameAsByte[i];
		}

		Comm.sendToServer(toServer);
	}

	// TODO test this
	/**
	 * Request
	 * 
	 * @param serverID
	 * @param roomID
	 * @return
	 * @throws IOException
	 */
	public void requestRoomDisplayName(int serverID, int roomID) throws IOException {

		byte[] toServer = new byte[9];
		toServer[0] = CommUtil.REQUEST_ROOM_DISPALY_NAME;

		byte[] idInByte = CommUtil.intToByteArr(serverID);
		byte[] idInByte2 = CommUtil.intToByteArr(roomID);

		for (int i = 0; i < idInByte.length; i++) {
			toServer[i + 1] = idInByte[i];
			toServer[i + 5] = idInByte2[i];
		}

		Comm.sendToServer(toServer);
	}

	public void requestAvatarURL(User user) throws IOException {
		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_USER_ICON;
		byte[] idInByte = CommUtil.intToByteArr(user.getId());

		for (int i = 0; i < idInByte.length; i++) {
			toServer[i + 1] = idInByte[i];
		}

		Comm.sendToServer(toServer);

	}

	/**
	 * Request the IP address of a user. If the client did not recieve data, it
	 * will return "FAIL"
	 * 
	 * @param User
	 *            The user to get the IP of
	 * @return The
	 * @throws IOException
	 */
	public String getUserIP(User User) throws IOException {
		IP = new String();
		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_USER_IP;

		byte[] usrID = CommUtil.intToByteArr(User.getId());
		for (int i = 0; i < usrID.length; i++) {
			toServer[i + 1] = usrID[i];

		}
		Comm.sendToServer(toServer);

		long oldTime = System.currentTimeMillis();

		while (IP.isEmpty() && System.currentTimeMillis() - oldTime < CommUtil.TIME_OUT_MILLI) {
			try {
				Thread.sleep(20);
			} catch (InterruptedException e) {
			}
		}

		if (IP.isEmpty()) {
			return "FAIL";
		}

		IP = IP.trim();
		return IP;

	}

	private void requestAllRoomsInServer(Server server) {
		byte[] toSend = new byte[5];
		toSend[0] = CommUtil.REQUEST_ALL_ROOM_ID;

		byte[] idInByte = CommUtil.intToByteArr(server.getId());
		for (int i = 0; i < idInByte.length; i++) {
			toSend[i + 1] = idInByte[i];
		}

		try {
			Comm.sendToServer(toSend);
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	public void requestServerDisplayName(Server server) {
		byte[] toSend = new byte[5];
		toSend[0] = CommUtil.REQUEST_SERVER_DISPLAY_NAME;

		byte[] idInByte = CommUtil.intToByteArr(server.getId());
		for (int i = 0; i < idInByte.length; i++) {
			toSend[i + 1] = idInByte[i];
		}

		try {
			Comm.sendToServer(toSend);
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	public Client getClient() {
		return this.mClient;
	}

	/**
	 * Send the status of the user to the server. This should be called whenever
	 * a status change is made from the user
	 */
	public void sendUserStatusChange() {
		byte status = (byte) 0;
		byte muted = (byte) 0;
		byte voice = (byte) 0;

		if (mClient.getIsMuted()) {
			muted = (byte) 1;
		}
		else {
			muted = (byte) 0;
		}

		if (mClient.isConnectedToVoice()) {
			voice = (byte) 1;
		}
		else {
			voice = (byte) 0;
		}

		if (mClient.getStatus() == CommUtil.UserStatus.OFFLINE) {
			status = (byte) 0;
		}
		else if (mClient.getStatus() == CommUtil.UserStatus.ONLINE) {
			status = (byte) 1;
		}
		else {
			status = (byte) 2;
		}

		byte[] toServer = new byte[4];

		toServer[0] = CommUtil.SEND_STAUTS_CHANGE;
		toServer[1] = status;
		toServer[2] = muted;
		toServer[3] = voice;

		try {
			Comm.sendToServer(toServer);
		} catch (IOException e) {

		}

	}
	
	public void requestUserStatus(User User){
		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_USER_STATUS;
		
		byte[] userID = CommUtil.intToByteArr(User.getId());
		
		for (int i = 0; i < userID.length; i++){
			toServer[1 + i] = userID[i];
		}
		
		try {
			Comm.sendToServer(toServer);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
		
		
		
	}

    public void changeLocation() {
        byte[] toSend = new byte[9];
        toSend[0] = CommUtil.SEND_ROOM_LOCATION_CHANGE;

        byte[] serverID = CommUtil.intToByteArr(mClient.getCurrentServerID());
        byte[] channelID = CommUtil.intToByteArr(mClient.getCurrentChannelID());

        for (int i = 0; i < serverID.length; i++){
            toSend[i+1] = serverID[i];
            toSend[i+5] = channelID[i];
        }

        try {
            Comm.sendToServer(toSend);
        } catch (IOException e) {
            e.printStackTrace();
        }

    }

}
