package com.jangle.communicate;

import java.io.*;
import java.net.*;

public class Client_Communicator {

	
	/**
	 * Creates a communication for the module, which can write to the write buffer, and read from the read buffer.  
	 */
	
	Socket Java_Socket;
	
	/**
	 * Object used to write to the socket
	 */
	PrintWriter Write;
	
	/**
	 * Object used to read form the socket
	 */
	BufferedReader Reader;
	
	
	
	public Client_Communicator(String Host, int port){
		
		try {
			Java_Socket= new Socket(Host, port);
		} catch (IOException e) {
			System.out.println("Could not initalize Socket for communicator");
			e.printStackTrace();
		}
		
		
		//Initialize PrintWriter to write to the output stream
		try {
			Write = new PrintWriter(Java_Socket.getOutputStream(), true);
		} catch (IOException e) {
			System.out.println("Could not initalize the writer");
			e.printStackTrace();
		}
		
		
		//Initialize buffer reader to read from the input stream
		try {
			Reader = new BufferedReader(new InputStreamReader(Java_Socket.getInputStream()));
		} catch (IOException e) {
			System.out.println("Could not initalize the reader");
			e.printStackTrace();
		}
		
		

	}
	
	
	
	/**
	 * Writes a string of data to the server
	 * @param Message The message to send to the server
	 */
	public void sendToServer(String Message){
		
		Write.println(Message);
		
	}
	
	/**
	 * Reads data from the server. This is a blocking call if there is nothing to read from the server
	 * @return The data read from the server
	 */
	public String readFromServer(){
		try {
			return Reader.readLine();
		} catch (IOException e) {
			System.out.println("Could not read from the server");
			e.printStackTrace();
		}
		return "Failed to read";
	}
	
	
	
	
	
}
