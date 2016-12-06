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

public class VoiceBroadcast implements Runnable {

	private ArrayList<VoiceChatSocket> connections;

	private AudioFormat format;
	private byte[] micData;
	private TargetDataLine microphone;
	private Client Cl;
	private ArrayList<User> Users;
	private Client_ParseData Parser;
	private DatagramSocket Send;

	private boolean sendAll;
	private int dataWidth;

	private int userID;
	private int port;

	public VoiceBroadcast(ArrayList<User> gUsers, AudioFormat gformat, Client gCl, int gport, Client_ParseData gParser,
			DatagramSocket gSend) {
		Users = gUsers;
		// connections = gConnections;
		sendAll = false;
		format = gformat;
		Cl = gCl;
		port = gport;
		Parser = gParser;
		Send = gSend;

		try {

			// init microphone
			DataLine.Info info = new DataLine.Info(TargetDataLine.class, format);
			microphone = (TargetDataLine) AudioSystem.getLine(info);
			micData = new byte[VoiceUtil.VOICE_DATA_SIZE];

		} catch (LineUnavailableException e) {
			e.printStackTrace();
		}
	}

	public void startBrodcast() {
		sendAll = true;
		try {
			startMicInput();
		} catch (LineUnavailableException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		Thread th = new Thread(this);
		th.start();
	}

	public void stopBrodcast() {
		sendAll = false;
		stopMic();
	}

	/**
	 * Start the input for the microphone. Input is whatever the default
	 * recording device of the operating system is
	 * 
	 * @throws LineUnavailableException
	 *             If the mic is not available. Could be because the microphone
	 *             was removed between now and object instantiation.
	 */

	private void startMicInput() throws LineUnavailableException {
		microphone.open(format);
		microphone.start();
	}

	/**
	 * Stop recording from mic.
	 */
	private void stopMic() {
		microphone.flush();
		microphone.stop();
		microphone.close();
	}

	@Override
	public void run() {
		int amount;
		DatagramPacket packet;

		while (sendAll) {
			if (Cl.getPushToTalk()) {
				int sum = 0;
				connections = new ArrayList<VoiceChatSocket>();

				amount = microphone.read(micData, 0, VoiceUtil.VOICE_DATA_SIZE);

				for (int i = 0; i < Users.size(); i++) {
					if (!Users.get(i).isChannel()) {
						if (Cl.getCurrentChannelID() == Users.get(i).getChannelID()
								&& Users.get(i).getIsMuted() == false && Users.get(i).getChannelID() != 0
								&& Users.get(i).getId() != Cl.getUserID()) {

							if (Users.get(i).getIP() == "" || Users.get(i).getIP() == "FAIL") {
								try {

									Users.get(i).setIP(Parser.getUserIP(Users.get(i)));

								} catch (IOException e) {
									// Happens if a communication error occurs.
								}
							}

							try {
								packet = new DatagramPacket(micData, micData.length,
										InetAddress.getByAddress(VoiceUtil.byteIP(Users.get(i).getIP())), port);
								// packet = new DatagramPacket(micData,
								// micData.length, InetAddress.getLocalHost(),
								// port);
							} catch (UnknownHostException e1) {
								continue;
							}

							try {
								Send.send(packet);
							} catch (IOException e1) {
							}

						}

					}
				}
			}

		}

	}

}
