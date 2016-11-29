package com.jangle.voice;

import java.io.IOException;
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

	private boolean sendAll;
	private int dataWidth;

	private int userID;
	private int port;

	public VoiceBroadcast(ArrayList<User> gUsers, AudioFormat gformat, Client gCl, int gport, Client_ParseData gParser) {
		Users = gUsers;
		// connections = gConnections;
		sendAll = false;
		format = gformat;
		Cl = gCl;
		port = gport;
		Parser = gParser;

		try {

			// init microphone
			DataLine.Info info = new DataLine.Info(TargetDataLine.class, format);
			microphone = (TargetDataLine) AudioSystem.getLine(info);
			micData = new byte[VoiceUtil.VOICE_DATA_BUFFER_SIZE];

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
		// TODO Auto-generated method stub

		while (sendAll) {

			connections = new ArrayList<VoiceChatSocket>();

			microphone.read(micData, 0, micData.length);
			/*
			 * This block is used if an external class/thread manages the
			 * connections ArrayList
			 * 
			 * byte[] toBrodcast = new byte[VoiceUtil.VOICE_DATA_BUFFER_SIZE];
			 * 
			 * for (int i = 0; i < connections.size(); i++){
			 * connections.get(i).sendVoice(micData); }
			 */

			for (int i = 0; i < Users.size(); i++) {
				Parser.requestUserStatus(Cl.getUsersArrayList().get(i));
				if (Cl.getCurrentChannelID() == Users.get(i).getChannelID() && Users.get(i).getIsMuted() == false) {
					/*
					 * NOTE, this does not care if the user wants to receive
					 * data, it will send it to users with the same channel ID.
					 * If the recieving user does not have their recieving
					 * enabled, the packet will get ignored on the reciever's
					 * end
					 */

					try {
						connections.add(new VoiceChatSocket(Users.get(i), port, Parser));
					} catch (IOException e) {
						// TODO Auto-generated catch block
						e.printStackTrace();
					}
					connections.get(i).sendVoice(micData);
				}

			}

			try {
				Thread.sleep(20);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}

	}

}
