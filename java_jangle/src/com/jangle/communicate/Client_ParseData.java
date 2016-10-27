package com.jangle.communicate;

import java.io.IOException;
import java.net.UnknownHostException;
import java.util.ArrayList;
import java.util.Arrays;

import com.jangle.client.*;
import com.jangle.communicate.CommUtil.*;

public class Client_ParseData implements IPARSER {

	private Client Cl;
	private Client_Communicator Comm;

	// Variables used when recieving data from the server. These are used as
	// temporary storage, and are not guaranted to hold any value
	private LoginResult loginResult;
	private int UserID;
	private String DisplayName;
	private int numMessagesRecieved;
	private String ServerDisplayName;
	private String RoomDisplayName;

	/**
	 * Create a parser object with no Client_Commmunicator attached to it.
	 * 
	 * @param Clie
	 *            The client object this communicator references
	 */
	public Client_ParseData(Client Clie) {
		this.Cl = Clie;

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
		this.Cl = Clie;
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
			Cl.addMessage(new Message(data));
			return;
		}

		if (data[0] == CommUtil.LOGIN_SUCCESS) {
			loginResult = LoginResult.SUCESS;
			UserID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
			return;
		}

		if (data[0] == CommUtil.LOGIN_FAIL) {
			loginResult = LoginResult.FAIL;
			return;
		}

		if (data[0] == CommUtil.CREATE_USER_FAIL) {
			loginResult = LoginResult.NAME_TAKEN;
			return;
		}

		if (data[0] == CommUtil.RECIEVE_DISPLAY_NAME) {
			DisplayName = new String(Arrays.copyOfRange(data, 5, data.length));
			return;
		}

		if (data[0] == CommUtil.RECIEVE_USERID) {

			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
			User tmp = new User("", id);
			// make sure the user is not in the list
			if (Cl.getUsers().contains(tmp) == false) {

				try {
					tmp.setDisplayName(this.requestDisplayName(tmp));
				} catch (IOException e) {
					e.printStackTrace();
				}

				Cl.addUser(tmp);
			}
		}

		if (data[0] == CommUtil.RECIEVE_SERVER_ID) {
			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
			//IMPLEMENT THIS LATER
		}
		
		if (data[0] == CommUtil.RECIEVE_SERVER_DISPLAY_NAME){
			ServerDisplayName = new String(Arrays.copyOfRange(data, 5, data.length));
		}
		
		if (data[0] == CommUtil.RECIEVE_ROOM_ID){
			int id = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
			// implement it later
		}
		
		if (data[0] == CommUtil.RECIEVE_ROOM_DISPLAY_NAME){
			RoomDisplayName = new String(Arrays.copyOfRange(data, 5, data.length));
			return;
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
	public LoginResult submitLogIn(String Username, String Password) throws IOException {

		byte[] data = new byte[20 + Password.length() + 1];
		loginResult = LoginResult.TIMEOUT;
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
		long startTime = System.currentTimeMillis();

		while ((loginResult == LoginResult.TIMEOUT)
				&& (System.currentTimeMillis() - startTime < CommUtil.TIME_OUT_MILLI)) {
		}

		if (loginResult == LoginResult.SUCESS) {
			Cl.setUserID(UserID);
		}

		UserID = 0;
		return loginResult;
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
	public LoginResult createUserInServer(String Username, String Password) throws IOException {

		byte[] data = new byte[20 + Password.length() + 1];
		loginResult = LoginResult.TIMEOUT;
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
		long startTime = System.currentTimeMillis();

		while ((loginResult == LoginResult.TIMEOUT)
				&& (System.currentTimeMillis() - startTime < CommUtil.TIME_OUT_MILLI)) {
		}

		if (loginResult == LoginResult.SUCESS) {
			Cl.setUserID(UserID);
		}

		UserID = 0;
		return loginResult;
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
	public void userIdTiedToServer(User user) throws IOException {

		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_ALL_SERVERID;

		byte[] tmp = CommUtil.intToByteArr(user.getId());

		for (int i = 0; i < tmp.length; i++) {
			toServer[i + 1] = tmp[i];
		}

		Comm.sendToServer(toServer);

	}

	// TODO SQL error?? preety sure my code is right
	public String requestDisplayName(User user) throws IOException {

		DisplayName = "";

		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_DISPLAY_NAME;

		byte[] idInByte = CommUtil.intToByteArr(user.getId());

		for (int i = 0; i < idInByte.length; i++) {
			toServer[i + 1] = idInByte[i];
		}

		Comm.sendToServer(toServer);
		long oldTime = System.currentTimeMillis();

		while (!DisplayName.equals("") && System.currentTimeMillis() - oldTime < 3000)
			;

		return DisplayName;

	}

	// TODO develop this one
	public User[] getUsersOnserver() {

		byte[] toServer = new byte[5];

		return null;
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

	// TODO test this
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

	// TODO Need to Test
	/**
	 * Request the Display name for a given serverID
	 * @param serverID the server you want the display name of
	 * @return the display name
	 * @throws IOException
	 */
	public String requestServerDisplayName(int serverID) throws IOException {

		ServerDisplayName = "";

		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_SERVER_DISPLAY_NAME;

		byte[] idInByte = CommUtil.intToByteArr(serverID);

		for (int i = 0; i < idInByte.length; i++) {
			toServer[i + 1] = idInByte[i];
		}

		Comm.sendToServer(toServer);
		long oldTime = System.currentTimeMillis();

		while (!ServerDisplayName.equals("") && System.currentTimeMillis() - oldTime < 3000)
			;

		return ServerDisplayName;

	}
	
	//TODO need to test this
	/**
	 * get a list of all of the room IDs in the room
	 * @param serverID
	 * @throws IOException 
	 */
	public void getRoomIDInServer(int serverID) throws IOException{
		byte[] toServer = new byte[4];
		byte[] nameAsByte = CommUtil.intToByteArr(serverID);
		toServer[0] = CommUtil.REQUEST_ALL_ROOM_ID;

		for (int i = 0; i < nameAsByte.length; i++) {
			toServer[i + 1] = nameAsByte[i];
		}

		Comm.sendToServer(toServer);
	}
	
	//TODO test this 
	/**
	 * Request 
	 * @param serverID
	 * @param roomID
	 * @return
	 * @throws IOException
	 */
	public String requestRoomDisplayName(int serverID, int roomID) throws IOException {

		RoomDisplayName = "";

		byte[] toServer = new byte[9];
		toServer[0] = CommUtil.REQUEST_ROOM_DISPALY_NAME;

		byte[] idInByte = CommUtil.intToByteArr(serverID);
		byte[] idInByte2 = CommUtil.intToByteArr(roomID);

		for (int i = 0; i < idInByte.length; i++) {
			toServer[i + 1] = idInByte[i];
			toServer[i + 5] = idInByte2[i];
		}

		Comm.sendToServer(toServer);
		long oldTime = System.currentTimeMillis();

		while (!RoomDisplayName.equals("") && System.currentTimeMillis() - oldTime < 3000)
			;

		return RoomDisplayName;

	}

	public Client getClient() {
		return this.Cl;
	}
}
