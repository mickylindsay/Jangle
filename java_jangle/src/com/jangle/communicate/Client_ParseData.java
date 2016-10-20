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

	// ints used when recieving data from the server. These are used as
	// temporary
	// storage, and are not guaranted to hold any value
	private LoginResult loginResult;
	private int UserID;
	private String DislayName;
	private int numMessagesRecieved;

	/**
	 * Create a parser object with no Client_Commmunicator attached to it.
	 * 
	 * @param Clie
	 *            The client object this communicator references
	 */
	public Client_ParseData(Client Clie) {
		Cl = Clie;

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
		Cl = Clie;
		Comm = new Client_Communicator(this, Host, port);
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
	 * userID of the client that is given at this parser instantiation will be set to the userID given by the server
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
	 * @param offSet Which block of 50 to 
	 * @return
	 * @throws IOException 
	 */
	public int request50MessagesWithOffset(int offSet) throws IOException{
		
		numMessagesRecieved = 0;
		
		byte[] test = new byte[2];
		test[0] = CommUtil.REQUEST_N_MESSAGES;
		test[1] = (byte) 0;
		
		Comm.sendToServer(test);
		
		return -1;
	}

	//Still broke yo, this aint done yet.
	public String requestDisplayName(User user) {

		DislayName = "";

		byte[] toServer = new byte[5];
		toServer[0] = CommUtil.REQUEST_DISPLAY_NAME;

		// Comm.sendToServer(Data);
		return "";

	}

	public Client getClient(){
		return this.Cl;
	}
}
