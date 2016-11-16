package com.jangle.voice;

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;
import java.net.UnknownHostException;
import java.util.ArrayList;

import javax.sound.sampled.AudioFormat;
import javax.sound.sampled.AudioSystem;
import javax.sound.sampled.DataLine;
import javax.sound.sampled.LineUnavailableException;
import javax.sound.sampled.SourceDataLine;
import javax.sound.sampled.TargetDataLine;

/**
 * Handles the creation of the voice chat. The recieving and playing to speakers are handled in this class.
 * So far, this class will handle the recieving. It can only handle one user at the moment.
 * 
 * @author Nathan Conroy
 *
 */
public class VoiceChat implements Runnable {

	private SourceDataLine speakers;

	private AudioFormat format;
	private int dataWidth;

	private int numChatWith;
	private ArrayList<VoiceChatSocket> connections;
	private DatagramSocket Recieving;
	private VoiceBroadcast Madden;
	
	private boolean isReceiving;

	private InetAddress Address;
	private int port;

	public VoiceChat(int gport) throws SocketException {
		format = new AudioFormat(8000.0f, 16, 1, true, true);
		try {
			// init speakers
			DataLine.Info dataLineInfo = new DataLine.Info(SourceDataLine.class, format);
			speakers = (SourceDataLine) AudioSystem.getLine(dataLineInfo);

		} catch (LineUnavailableException e) {
			e.printStackTrace();
		}

		numChatWith = 0;
		isReceiving = false;
		connections = new ArrayList<VoiceChatSocket>();
		port = gport;

		try {
			Address = InetAddress.getLocalHost();
		} catch (UnknownHostException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
		Madden = new VoiceBroadcast(connections, format);
		Recieving = new DatagramSocket(gport);
		dataWidth = Madden.getDataWidth();

	}

	/**
	 * Add a user. This adds a VoiceChatSocket. for Testing, you can put in
	 * local host, and hear yourself
	 * 
	 * @param IP
	 *            IP of the user.
	 */
	public void addUserToChat(String IP) {
		try {
			connections.add(new VoiceChatSocket(IP, port, dataWidth));
		} catch (IOException e) {
			e.printStackTrace();
		}
		numChatWith++;
	}

	public void closeAllConctions() {

	}
	
	public void startBrodcast(){
		try {
			Madden.startMicInput();
		} catch (LineUnavailableException e) {
			System.out.println("Failed to start mic");
			e.printStackTrace();
		}
		Madden.brodcastToAll();
	}

	/**
	 * Start the output of audio. Will play the sound to the default device of
	 * the operating systems
	 * 
	 * @throws LineUnavailableException
	 *             If the speaker is not instantiation. Could be due to the
	 *             speaker was removed from now and this object's instantiation
	 */
	public void startSpeakers() throws LineUnavailableException {
		speakers.open(format);
		speakers.start();
	}

	/**
	 * Stop playing to the speakers.
	 */
	public void stopSpeakers() {
		speakers.drain();
		speakers.close();
	}

	public void recieveData() {
		isReceiving = true;
		Thread th = new Thread(this);
		th.start();
	}
	
	public void stopRecieve(){
		isReceiving = false;
	}

	@Override
	public void run() {

		// TODO with code below in a thead from main, this works. Need to put in
		// differnet code from the voice part. Think about a differnet class.
		// Also remove SYSO
		while (true) {
			byte[] data = new byte[1024];
			DatagramPacket packet = new DatagramPacket(data, data.length);
			try {
				Recieving.receive(packet);
			} catch (IOException e) {
				e.printStackTrace();
			}

			speakers.write(data, 0, data.length);

			try {
				Thread.sleep(20);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}

}