package com.jangle.communicate;

import java.io.IOException;
import java.net.UnknownHostException;
import java.util.Arrays;

import com.jangle.client.*;
import com.jangle.communicate.CommUtil.*;

public class Client_ParseData implements IPARSER {

	private Client Cl;
	private Client_Communicator Comm;
	
	private LoginResult loginResult;
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
			return;
		}
		
		if (data[0] == CommUtil.LOGIN_SUCCESS){
			loginResult = LoginResult.SUCESS;
			UserID = CommUtil.byteToInt(Arrays.copyOfRange(data, 1, data.length));
			return;
		}
		
		if (data[0] == CommUtil.LOGIN_FAIL){
			loginResult = LoginResult.FAIL;
			return;
		}
		
		if (data[0] == CommUtil.CREATE_USER_FAIL){
			loginResult = LoginResult.FAIL;
		}
		
		
	}
	
	
	

	 /**
	  * 
	  * @param user user object for the user logging in. If succes, the user ID will be places in this user.
	  * @param Password Password for the user
	  * @return
	  * @throws IOException
	  */
	public LoginResult submitLogIn(User user, String Password) throws IOException{
		
		byte[] data = new byte[user.getUserName().length() + Password.length() + 1];
		data[0] = CommUtil.LOGIN;
		loginResult = LoginResult.TIMEOUT;
		
		Comm.sendToServer(data);
		
		//Sleep the thread, so we can give time to the thread listing to the receiver
		//time to update values.
		try {
			Thread.sleep(2000);
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
		if (loginResult == LoginResult.SUCESS){
			user.setId(UserID);
		}
		return loginResult;
	}
	
	
	/**
	 * Submit request to create a new user
	 * @param Username username of the new user
	 * @param Password password for the new user
	 * @return 
	 * @throws IOException 
	 */
	public LoginResult createUserInServer(User user, String Password) throws IOException{
		byte[] data = new byte[user.getUserName().length() + Password.length() + 1];
		data[0] = CommUtil.CREATE_USER;
		loginResult = LoginResult.TIMEOUT;
		
		Comm.sendToServer(data);
		
		//Sleep the thread, so we can give time to the thread listing to the receiver
		//time to update values.
		try {
			Thread.sleep(2000);
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
		if (loginResult == LoginResult.SUCESS){
			user.setId(UserID);
		}
		return loginResult;
	}


}
