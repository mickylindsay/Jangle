package com.jangle.communicate;

import java.io.*;
import java.net.*;
import com.jangle.communicate.CommUtil;

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
		byte[] dataFromServer = new byte[1024];
		byte[] dataStore;
		byte[] dataToRet = new byte[0];
		byte[] tmp = new byte[4];
		int amount = 0;
		int bytesToRead = 0;

		try {
			amount += Reader.read(dataFromServer);

		} catch (SocketTimeoutException ste) {
			System.out.println("no");
		}
		
		tmp[0] = dataFromServer[0];
		tmp[1] = dataFromServer[1];
		tmp[2] = dataFromServer[2];
		tmp[3] = dataFromServer[3];
		
		bytesToRead = CommUtil.byteToInt(tmp);
		
		if (bytesToRead < amount) {
			byte [] ret = new byte[dataFromServer.length - 4];
			for (int i = 0; i < ret.length; i++){
				ret[i] = dataFromServer[i + 4];
			}
			return ret;
		}
		
		while (true) {
			
			try {
				amount += Reader.read(dataFromServer);

			} catch (SocketTimeoutException ste) {
				System.out.println("no");
			}
			
			dataStore = dataToRet.clone();
			dataToRet = new byte[dataStore.length + dataFromServer.length];
			
			for (int i = 0; i < dataStore.length; i++){
				dataToRet[i] = dataStore[i];
			}
			
			for (int i = 0; i < dataFromServer.length; i++){
				dataToRet[dataStore.length + i] = dataFromServer[i];
			}
			
			if (bytesToRead < amount){
				byte[] ret = new byte[dataToRet.length - 4];
				for (int i = 0; i < ret.length; i++){
					ret[i] = dataToRet[i + 4];
				}
			}
			

		}
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
			try {
				Thread.sleep(50);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}

	}
}
