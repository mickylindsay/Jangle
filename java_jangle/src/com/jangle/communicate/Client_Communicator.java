package com.jangle.communicate;

import java.io.*;
import java.net.*;

public class Client_Communicator {

	/**
	 * Creates a communication for the module, which can write to the write
	 * buffer, and read from the read buffer.
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

	public Client_Communicator(String Host, int port) throws UnknownHostException, IOException {

		Java_Socket = new Socket(Host, port);

		// Initialize PrintWriter to write to the output stream

		Write = new PrintWriter(Java_Socket.getOutputStream(), true);

		// Initialize buffer reader to read from the input stream
		Reader = new BufferedReader(new InputStreamReader(Java_Socket.getInputStream()));

	}

	/**
	 * Writes a string of data to the server
	 * 
	 * @param Message
	 *            The message to send to the server
	 */
	public void sendToServer(byte[] Message) {

		Write.println(Message);

	}

	/**
	 * Reads data from the server. This is a blocking call if there is nothing
	 * to read from the server
	 * 
	 * @return The data read from the server
	 * @throws IOException 
	 */
	public String readFromServer() throws IOException {
		 return Reader.readLine();
	}

}
