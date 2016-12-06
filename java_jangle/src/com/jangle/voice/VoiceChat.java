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
	private int port;
	private int userID;

	public VoiceChat(int gport, boolean speak, Client gCl, Client_ParseData gParser) throws SocketException {
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
		Recieving = new DatagramSocket(gport);
		Recieving.setReceiveBufferSize(3);
		
		Madden = new VoiceBroadcast(Users, format, Cl, port, Parser, Recieving);

		// If speak is true, the user wants to start speaking right away
		if (speak) {
			connectToVoice();
			startBrodcast();
		}

	}

	/**
	 * Want connect the user to the voice chat. Does not start broadcasting
	 * voice. However, data that is sent from users in the voice chat will play
	 * though the device's default audio device
	 * 
	 * To start sending voice to other users, call the function StartBrodcast();
	 */
	public void connectToVoice() {
		if (!Cl.isConnectedToVoice()) {

			// Start speakers
			try {
				startSpeakers();
			} catch (LineUnavailableException e) {
				// Speakers are not ready to broadcast to.
			}

			// start recieving data
			recieveData();
			Cl.setConnectedToVocie(true);
			Parser.sendUserStatusChange();

		}
	}

	/**
	 * Disconnect the user from Voice chat. The user does not want to be part of
	 * the voice chat
	 */
	public void disconnectFromVoice() {
		stopRecieve();
		System.out.println("Stopped Recieveing");
		connections.clear();
		stopSpeakers();
		
		if (Cl.getBroadcastStatus()){
			endBrodcast();
		}
		Cl.setConnectedToVocie(false);
		Parser.sendUserStatusChange();
	}

	/**
	 * Start sending voice to other users. You can only send voice data to other
	 * users if you are connected to voice chat
	 */
	public void startBrodcast() {
		if (!Cl.getBroadcastStatus() && Cl.isConnectedToVoice()) {
			Madden.startBrodcast();
			Cl.setBroadcastStatus(true);
			Parser.sendUserStatusChange();
		}
	}

	/**
	 * Stop sending voice to other users. However, user is still connected to
	 * the voice chat, and will still be receiving voice data
	 */
	public void endBrodcast() {
		Madden.stopBrodcast();
		Cl.setBroadcastStatus(false);
		Parser.sendUserStatusChange();
	}

	/*
	 * IS NOT USED. DO NOT USE THIS /** Add a user. This adds a VoiceChatSocket.
	 * for Testing, you can put in local host, and hear yourself
	 * 
	 * @param IP IP of the user.
	 * 
	 * @throws IOException
	 * 
	 * @throws UnknownHostException
	 * 
	 * private void addUserToChat(User gUser) throws UnknownHostException,
	 * IOException { connections.add(new VoiceChatSocket(gUser, port, Parser));
	 * }
	 */

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

	/**
	 * Calculate the number of users in the same channel that are talking
	 * Assuming that the user array list is all of the users in the same server
	 * 
	 * @return
	 */
	private int numUsersInSameChannel() {
		int ret = 0;

		for (int i = 0; i < Cl.getUsersArrayList().size(); i++) {
			if (Cl.getUsersArrayList().get(i).getChannelID() == Cl.getCurrentChannelID()
					&& Cl.getUsersArrayList().get(i).getVoiceStatus()) {
				ret += 1;
			}
		}

		return ret;
	}

	@Override
	public void run() {
		byte[] toSpeaker = new byte[VoiceUtil.VOICE_DATA_SIZE];
		byte[] amb = new byte[4];
		DatagramPacket packet = new DatagramPacket(toSpeaker, toSpeaker.length);
		int loop = 1;
		int numUsers = 0;
		int amountRead = 0;

		for (int i = 0; i < toSpeaker.length; i++) {
			toSpeaker[i] = 0;
		}
		
		
		while (isReceiving) {

			try {
				Recieving.receive(packet);
				toSpeaker = packet.getData();
			} catch (IOException e) {
				//stuff
			}
			/**
			amb[0] = data[0];
			amb[1] = data[1];
			amb[2] = data[2];
			amb[3] = data[3];
			
			amountRead += VoiceUtil.byteToInt(amb);
			*/
			numUsers = numUsersInSameChannel();
			if (numUsers == 0) {
				continue;
			}
			
			speakers.write(toSpeaker, 0, toSpeaker.length);
			
			/*
			if (loop % numUsers == 0) {
				loop = 0;

				for (int i = 0; i < toSpeaker.length; i++) {
					//toSpeaker[i] = (byte) ((data[i + 4] + toSpeaker[i]) >> 1);
					toSpeaker[i] = data[i];
				} 

				speakers.write(toSpeaker, 0, amountRead);
				toSpeaker = new byte[VoiceUtil.VOICE_DATA_SIZE];
				amountRead = 0;
			}
			else {
				for (int i = 0; i < toSpeaker.length; i++) {
					toSpeaker[i] = data[i + 4];
				}
			}
			loop += 1;
			*/

		}
	}

}
