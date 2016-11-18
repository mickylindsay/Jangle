package com.jangle.voice;

import java.util.ArrayList;

import javax.sound.sampled.AudioFormat;
import javax.sound.sampled.AudioSystem;
import javax.sound.sampled.DataLine;
import javax.sound.sampled.LineUnavailableException;
import javax.sound.sampled.SourceDataLine;
import javax.sound.sampled.TargetDataLine;
import com.jangle.client.User;

public class VoiceBroadcast implements Runnable {

	private ArrayList<VoiceChatSocket> connections;

	private AudioFormat format;
	private byte[] micData;
	private TargetDataLine microphone;

	private boolean sendAll;
	private int dataWidth;
	
	private int userID;

	public VoiceBroadcast(ArrayList<VoiceChatSocket> gConnections, AudioFormat gformat, int gUser) {
		connections = gConnections;
		sendAll = false;
		format = gformat;
		userID = gUser;

		try {

			// init microphone
			DataLine.Info info = new DataLine.Info(TargetDataLine.class, format);
			microphone = (TargetDataLine) AudioSystem.getLine(info);
			micData = new byte[VoiceUtil.VOICE_DATA_BUFFER_SIZE];

		} catch (LineUnavailableException e) {
			e.printStackTrace();
		}

	}
	/**
	 * Start the input for the microphone. Input is whatever the default
	 * recording device of the operating system is
	 * 
	 * @throws LineUnavailableException
	 *             If the mic is not available. Could be because the microphone
	 *             was removed between now and object instantiation.
	 */

	public void startMicInput() throws LineUnavailableException {
		microphone.open(format);
		microphone.start();
	}
	
	/**
	 * Stop recording from mic.
	 */
	public void stopMic() {
		microphone.close();
		stopBrodcast();
	}

	public void brodcastToAll() {
		sendAll = true;
		Thread th = new Thread(this);
		th.start();
	}

	public void stopBrodcast() {
		sendAll = false;
	}

	@Override
	public void run() {
		// TODO Auto-generated method stub

		while (sendAll) {
			microphone.read(micData, 0, micData.length);
			
			byte[] toBrodcast = new byte[VoiceUtil.VOICE_DATA_BUFFER_SIZE + 4];
			
			
			connections.get(0).sendVoice(micData);
			try {
				Thread.sleep(20);
			} catch (InterruptedException e) {
				// TODO Auto-generated catch block
				e.printStackTrace();
			}
		}

	}

}
