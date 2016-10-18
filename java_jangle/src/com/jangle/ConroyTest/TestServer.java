package com.jangle.ConroyTest;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;

public class TestServer implements Runnable  {
	
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
	
	public TestServer( int port){
		try {
			Comm = new ServerSocket(port);
		} catch (IOException e) {
			
			System.out.println("\n\nCannot create test server\n\n");
		}

		// Initialize PrintWriter to write to the output stream
		
		Thread t = new Thread(this);
		t.start();
		
	}
	
	public void sendToClient(byte[] data){
		try {
			Write.write(data);
		} catch (IOException e) {
			System.out.println("\n\nCannot send data to client\n\n");
		}
		
		System.out.println("Sent data to client");
	}
	
	public void readFromClient(){
		byte[] dataFromClient = new byte[1024];
		
		try {
			System.out.println(Reader.read(dataFromClient));
		} catch (IOException e) {
			System.out.println("\n\nCANT READ DATA FROM CLIENT YO\n\n");
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
		while (true){
			readFromClient();
		}
		
	}

}
