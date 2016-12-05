package com.jangle.communicate;

import java.io.*;
import java.net.*;
import com.jangle.communicate.CommUtil;

public class Client_Communicator implements Runnable {

	/**
	 * Creates a communication for communication with the server, which can write to the write
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

    private boolean done;

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
        this.done = false;

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
		byte[] toServer = new byte[Data.length + 4];
		byte[] tmp = CommUtil.intToByteArr(Data.length);


		for (int i = 0; i < tmp.length; i++) {
			toServer[i] = tmp[i];
		}

		for (int i = 0; i < Data.length; i++) {
			toServer[4 + i] = Data[i];
		}

		Write.write(toServer);

	}

	/**
	 * Reads data from the server. This is a blocking call if there is nothing
	 * to read from the server
	 * 
	 * @return The data read from the server
	 * @throws IOException
	 */
	private byte[] readFromServer() throws IOException {
		byte[] tmp = new byte[4];
		int amount = 0;
		int bytesToRead = 0;

		try {
			amount = Reader.read(tmp);

		} catch (SocketTimeoutException ste) {
			System.out.println("no");
		}

		byte[] dataFromServer = new byte[CommUtil.byteToInt(tmp)];

		try {
			amount = Reader.read(dataFromServer);

		} catch (SocketTimeoutException ste) {
			System.out.println("no");
		}

		
		
		return dataFromServer;

	}

    public void endThread() {
        this.done = true;
    }

	@Override
	public void run() {

		while (!done) {
			
			try {
				
				byte[] fromServer = readFromServer();
				
				if (fromServer != null) {
					
					Parser.parseData(fromServer);
				}

				fromServer = null;

			} catch (IOException e1) {
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
