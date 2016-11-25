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

import com.jangle.client.Client;
import com.jangle.client.User;
import com.jangle.communicate.Client_ParseData;

/**
 * Handles the creation of the voice chat. The recieving and playing to speakers
 * are handled in this class. So far, this class will handle the recieving. It
 * can only handle one user at the moment.
 * 
 * @author Nathan Conroy
 *
 */
public class VoiceChat implements Runnable {

	private SourceDataLine speakers;

	private AudioFormat format;
	private ArrayList<VoiceChatSocket> connections;
	private DatagramSocket Recieving;
	private VoiceBroadcast Madden;
	private Client Cl;
	private Client_ParseData Parser;

	private ArrayList<User> Users;

	private boolean isReceiving;

	private InetAddress Address;
	private int port;
	private int userID;

	public VoiceChat(int gport, boolean speak, Client gCl,  Client_ParseData gParser) throws SocketException {
		format = VoiceUtil.genFormat();
		try {
			// init speakers
			DataLine.Info dataLineInfo = new DataLine.Info(SourceDataLine.class, format);
			speakers = (SourceDataLine) AudioSystem.getLine(dataLineInfo);

		} catch (LineUnavailableException e) {
			e.printStackTrace();
		}

		isReceiving = false;
		connections = new ArrayList<VoiceChatSocket>();
		port = gport;
		Cl = gCl;
		Users = Cl.getUsersArrayList();
		Parser = gParser;
		

		try {
			Address = InetAddress.getLocalHost();
		} catch (UnknownHostException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}

		Madden = new VoiceBroadcast(Users, format, Cl, port, Parser);
		Recieving = new DatagramSocket(gport);
		
		
		//If speak is true, the user wants to start speaking right away
		if (speak){
			connectToVoice();
			startBrodcast();
		}

	}

	/**
	 * Want connect the user to the voice chat. Does not start broadcasting voice.
	 * However, data that is sent from users in the voice chat will play though
	 * the device's default audio device
	 * 
	 * To start sending voice to other users, call the function StartBrodcast();
	 */
	public void connectToVoice() {
		//Start speakers
		try {
			startSpeakers();
		} catch (LineUnavailableException e) {
			//Speakers are not ready to broadcast to.
		}
		
		//start recieving data
		recieveData();
		
	}

	public void disconnectFromVoice() {
		connections.clear();
		stopSpeakers();
		stopRecieve();
		endBrodcast();
	}

	public void startBrodcast() {
		Madden.startBrodcast();
	}

	public void endBrodcast() {
		Madden.stopBrodcast();
	}

	/**
	 * Add a user. This adds a VoiceChatSocket. for Testing, you can put in
	 * local host, and hear yourself
	 * 
	 * @param IP
	 *            IP of the user.
	 * @throws IOException
	 * @throws UnknownHostException
	 */
	private void addUserToChat(User gUser) throws UnknownHostException, IOException {
		connections.add(new VoiceChatSocket(gUser, port, Parser));
	}

	/**
	 * Start the output of audio. Will play the sound to the default device of
	 * the operating systems
	 * 
	 * @throws LineUnavailableException
	 *             If the speaker is not instantiation. Could be due to the
	 *             speaker was removed from now and this object's instantiation
	 */
	private void startSpeakers() throws LineUnavailableException {
		speakers.open(format);
		speakers.start();
	}

	/**
	 * Stop playing to the speakers.
	 */
	private void stopSpeakers() {
		speakers.drain();
		speakers.close();
	}

	private void recieveData() {
		isReceiving = true;
		Thread th = new Thread(this);
		th.start();
	}

	private void stopRecieve() {
		isReceiving = false;
	}

	@Override
	public void run() {
		while (isReceiving) {
			byte[] data = new byte[VoiceUtil.VOICE_DATA_BUFFER_SIZE];
			byte[] toSpeaker = new byte[VoiceUtil.VOICE_DATA_BUFFER_SIZE];
			DatagramPacket packet = new DatagramPacket(data, data.length);
			int loop = 0;
			
			try {
				Recieving.receive(packet);
			} catch (IOException e) {
				e.printStackTrace();
			}

			if (loop % Cl.getUsers().size() == 0){
				
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
