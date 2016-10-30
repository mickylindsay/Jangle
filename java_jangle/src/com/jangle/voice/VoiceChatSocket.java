package com.jangle.voice;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.Socket;
import java.net.UnknownHostException;

import com.jangle.client.User;

/**
 * Class that acts as a wrapper for a socket, so is easier to manage for
 * VoiceChat
 * 
 * @author Nathan Conroy
 *
 */
public class VoiceChatSocket {

	/**
	 * Creates a communication for communication with the server, which can
	 * write to the write buffer, and read from the read buffer.
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
	 * The width of the buffer for mic data
	 */
	private int micDataWidth;

	private User User;

	private String adress;

	/**
	 * Create socket to communicate with
	 * 
	 * @param Adress
	 * @param port
	 * @throws UnknownHostException
	 * @throws IOException
	 */
	public VoiceChatSocket(String Adress, int port, int micDataWidth, User gUser)
			throws UnknownHostException, IOException {
		adress = Adress;
		User = gUser;

		Java_Socket = new Socket(Adress, port);

		// Initialize PrintWriter to write to the output stream
		Write = Java_Socket.getOutputStream();

		// Initialize buffer reader to read from the input stream
		Reader = Java_Socket.getInputStream();

		this.micDataWidth = micDataWidth;
	}

	public byte[] recieveVoice() throws IOException {

		byte[] ret = new byte[micDataWidth];
		Reader.read(ret);
		return ret;
	}

	public void sendMic(byte[] data) throws IOException {
		Write.write(data);
	}
}
