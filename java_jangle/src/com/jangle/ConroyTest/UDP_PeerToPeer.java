package com.jangle.ConroyTest;

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;
import java.net.UnknownHostException;


public class UDP_PeerToPeer {

	private DatagramSocket socket;
	private InetAddress Address;
	private int port;

	public UDP_PeerToPeer(String gAddress, int gport) throws SocketException, UnknownHostException {
		port = gport;
		socket = new DatagramSocket(port);
		Address = InetAddress.getByName(gAddress);
		
	}

	public void sendData(byte[] data) throws IOException {
		DatagramPacket packet = new DatagramPacket(data, data.length, Address, port);
		socket.send(packet);
	}
	
	public byte[] recieveByte() throws IOException{
		byte[] data = new byte[4];
		//TODO figure out the byte width for here
		DatagramPacket packet = new DatagramPacket(data, data.length);
		socket.receive(packet);
		return data;
	}
	
}
