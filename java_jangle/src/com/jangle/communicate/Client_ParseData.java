package com.jangle.communicate;

import java.io.IOException;
import java.net.UnknownHostException;
import java.util.Arrays;

import com.jangle.client.*;
import com.jangle.communicate.CommUtil.*;

public class Client_ParseData implements IPARSER {

	private Client Cl;
	private Client_Communicator Comm;
	
	private int loginResult;
	private int UserID;
	
	
	

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

		//when ready to send opcode
		//Comm.sendToServer(generateMessage(mess)); when opcode is ready;
		Comm.sendToServer(mess.getByteArray());
	}

	
	/**
	 * Figure out what the data that was received is.
	 * 
	 * @param data the character array to parse, and figure out what it is
	 */
	public void parseData(byte[] data) {

		if (data[0] == CommUtil.MESSAGE_FROM_SERVER){
			Cl.addMessage(new Message (data));
		}
		
		if (data[1] == CommUtil.LOGIN_SUCCESS){
			loginResult = 1;
			UserID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
		}
		
		if (data[2] == CommUtil.LOGIN_FAIL){
			loginResult = 0;
		}
	}
	
	
	
	/**
	 * Submit username and password for log in. If the login is true, the userID for that user
	 * can be retrieved by calling getUserIDLastLogin()
	 * @param Username the username to log in with
	 * @param Password password to log in with
	 * @param UserId If the login is a success, the UserID will be put here
	 * @return 0 if login failed. 1 if success. -1 if there was a time out
	 * @throws IOException 
	 */
	public int submitLogIn(String Username, String Password, String UserId) throws IOException{
		
		Username = Username.trim();
		Password = Password.trim();
		byte[] data = new byte[Username.length() + Password.length() + 1];
		UserID = 0;
		
		
		data[0] = (byte) 2;
		loginResult = -1;
		
		Comm.sendToServer(data);
		
		//Sleep the thread, so we can give time to the thread listing to the receiver
		//time to update values.
		try {
			Thread.sleep(2000);
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		return loginResult;
		
		
	}
	
	public int getUserIDLastLogin(){
		return UserID;
	}
	
	/**
	 * Submit request to create a new user
	 * @param Username username of the new user
	 * @param Password password for the new user
	 * @return 0 if login failed. 1 if success. -1 if there was a time out
	 */
	public int createUser(String Username, String Password){
		return 0;
	}


}
