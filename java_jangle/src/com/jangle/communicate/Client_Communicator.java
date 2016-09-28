package com.jangle.communicate;

import java.io.*;
import java.net.*;
import com.jangle.communicate.Comm_CONSTANTS;
//add import for message 

public class Client_Communicator implements Runnable{

	/**
	 * Creates a communication for the module, which can write to the write
	 * buffer, and read from the read buffer.
	 */

	Socket Java_Socket;

	/**
	 * Object used to write to the socket
	 */
	OutputStream Write;

	/**
	 * Object used to read form the socket
	 */
	BufferedReader Reader;
	
	Client_ParseData parser;
	
	//JavaUI UI;

	public Client_Communicator(String Host, int port) throws UnknownHostException, IOException {

		Java_Socket = new Socket(Host, port);
		Java_Socket.setSendBufferSize(1024);
		Java_Socket.setReceiveBufferSize(1024);

		// Initialize PrintWriter to write to the output stream

		Write = Java_Socket.getOutputStream();
		
		//create parser object
		parser = new Client_ParseData();

		// Initialize buffer reader to read from the input stream
		Reader = new BufferedReader(new InputStreamReader(Java_Socket.getInputStream()));
		Thread t = new Thread(this);
		t.start();
		
		
		//UI = givenUI;
		

	}
	
	public void sendMessage(Message mess, int serverID, int channedID) {
		// TODO Auto-generated method stub
		
	}
	



	

	/**
	 * Writes a string of data to the server
	 * 
	 * @param Message
	 *            The message to send to the server
	 * @throws IOException 
	 */
	private void sendToServer(byte[] Message) throws IOException {

		String s = String.valueOf(Message);
		
		Write.write(Message);
		
		System.out.println(Message);

	}

	/**
	 * Reads data from the server. This is a blocking call if there is nothing
	 * to read from the server
	 * 
	 * @return The data read from the server
	 * @throws IOException 
	 */
	private String readFromServer() throws IOException {
		 return Reader.readLine();
	}

	@Override
	public void run() {
		// TODO Auto-generated method stub
		
		while (true){
			
			try {
				if (readFromServer() == null){
					
				}
				else{
					//UI.sendMessage()
				}
			} catch (IOException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
			
			
			try {
				Thread.sleep(200);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}
		
	}






	@Override
	public void sendMessage(Message mess, int serverID, int channedID) {
		// TODO Auto-generated method stub
		
	}



}
