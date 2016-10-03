package com.jangle.communicate;

import java.io.*;
import java.net.*;
import com.jangle.communicate.Comm_CONSTANTS;
//add import for message 

public class Client_Communicator implements Runnable {

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
	InputStream Reader;

	/**
	 * Parser object this Client_Communicator calls on to parse data
	 */
	Client_ParseData Parser;

	/**
	 * 
	 * @param gParser
	 *            The parser the Clien_Communicator calls on to parse data
	 * @param Host
	 *            The IP address of the server
	 * @param port
	 *            port to communicate though with the server
	 * @throws UnknownHostException
	 * @throws IOException
	 */
	public Client_Communicator(Client_ParseData gParser, String Host, int port)
			throws UnknownHostException, IOException {

		Java_Socket = new Socket(Host, port);
		Parser = gParser;

		// Initialize PrintWriter to write to the output stream

		Write = Java_Socket.getOutputStream();

		// Initialize buffer reader to read from the input stream
		Reader = Java_Socket.getInputStream();
		// Java_Socket.setSoTimeout(5);
		Thread t = new Thread(this);
		t.start();
	}

	/**
	 * Writes a string of data to the server
	 * 
	 * @param Data
	 *            the data to send to the server
	 * @throws IOException
	 */
	public void sendToServer(byte[] Data) throws IOException {
		Write.write(Data);

	}

	/**
	 * Reads data from the server. This is a blocking call if there is nothing
	 * to read from the server
	 * 
	 * @return The data read from the server
	 * @throws IOException
	 */
	private byte[] readFromServer() throws IOException {
		byte[] tmp = new byte[1024];
		int amount;
		try {
			amount = Reader.read(tmp);
			System.out.println(amount);
			return tmp;

		} catch (SocketTimeoutException ste) {
			System.out.println("no");
		}
		return null;
	}

	@Override
	public void run() {

		while (true) {
			byte[] tmp = new byte[1024];
			try {
				tmp = readFromServer();
				if (tmp != null) {
					Parser.parseData(tmp);
				}

				tmp = null;

			} catch (IOException e1) {
				// TODO Auto-generated catch block
				e1.printStackTrace();
			}
		}

	}
}
