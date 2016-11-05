package com.jangle.voice;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
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
public class VoiceChatSocket implements Runnable {

	private DatagramSocket socket;
	private InetAddress Address;
	private int port;
	
	private byte[] Data;
	
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
	public VoiceChatSocket(String gAddress, int gport, int gmicDataWidth)
			throws UnknownHostException, IOException {
		port = gport;
		socket = new DatagramSocket();
		Address = InetAddress.getByName(gAddress);

		this.micDataWidth = gmicDataWidth;
		
		
	}
	
	public void sendVoice(byte[] data){
		Data = data;
		Thread th = new Thread(this);
		th.start();
	}
	

	@Override
	public void run() {
		// TODO Auto-generated method stub
		
		DatagramPacket packet = new DatagramPacket(Data, micDataWidth, Address, port);
		try {
			socket.send(packet);
		} catch (IOException e) {
			System.out.println("fail");
			e.printStackTrace();
		}
		
	}

}
