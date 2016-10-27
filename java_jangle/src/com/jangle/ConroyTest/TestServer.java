package com.jangle.ConroyTest;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;
import java.net.SocketTimeoutException;

import com.jangle.communicate.CommUtil;

public class TestServer implements Runnable {

	private ServerSocket Comm;

	private Socket ServerSock;

	/**
	 * Object used to write to the socket
	 */
	OutputStream Write;

	/**
	 * Object used to read form the socket
	 */
	InputStream Reader;

	public TestServer(int port) {
		try {
			Comm = new ServerSocket(port);
		} catch (IOException e) {

			System.out.println("\n\nCannot create test server\n\n");
		}

		// Initialize PrintWriter to write to the output stream

		Thread t = new Thread(this);
		t.start();

	}

	/**
	 * Does not have long data built yet
	 * 
	 * @param data
	 */
	public void sendToClient(byte[] data) {

		byte[] data2 = new byte[data.length + 4];
		int length = data.length;

		byte[] size = CommUtil.intToByteArr(length);

		data2[0] = size[0];
		data2[1] = size[1];
		data2[2] = size[2];
		data2[3] = size[3];

		for (int i = 0; i < data.length; i++) {
			data2[i + 4] = data[i];
		}

		System.out.println();
		try {
			Write.write(data2);
		} catch (IOException e) {
			System.out.println("\n\nCannot send data to client\n\n");
		}

		System.out.println("Sent " + data2.length + " bytes to client");
	}

	public void readFromClient() throws IOException {

		byte[] tmp = new byte[4];
		int amount = 0;
		int bytesToRead = 0;

		try {
			amount = Reader.read(tmp);
			System.out.println(amount);

		} catch (SocketTimeoutException ste) {
			System.out.println("no");
		}

		System.out.println("Client will read " + CommUtil.byteToInt(tmp));

		byte[] dataFromServer = new byte[CommUtil.byteToInt(tmp)];

		try {
			amount = Reader.read(dataFromServer);
			System.out.println(amount + "\n");

		} catch (SocketTimeoutException ste) {
			System.out.println("no");
		}

		if (dataFromServer != null) {

			System.out.println("stuff from the client\n");
			for (int i = 0; i < dataFromServer.length; i++) {
				System.out.print(dataFromServer[i] + " ");
			}
		}
	}

	@Override
	public void run() {
		System.out.println("Wating for user");
		try {
			ServerSock = Comm.accept();
		} catch (IOException e) {
			System.out.println("\n\nCannot accept user\n\n");
		}

		try {
			Write = ServerSock.getOutputStream();
		} catch (IOException e) {
			System.out.println("\n\nCannot create writer\n\n");
		}

		// Initialize buffer reader to read from the input stream
		try {
			Reader = ServerSock.getInputStream();
		} catch (IOException e) {
			System.out.println("\n\nCannot create reader\n\n");
		}

		System.out.println("Test server initalized");
		System.out.println("User Connected\n\n");
		while (true) {
			try {
				readFromClient();
			} catch (IOException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}

	}

}
