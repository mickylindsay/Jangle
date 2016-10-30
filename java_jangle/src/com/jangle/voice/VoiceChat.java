package com.jangle.voice;

import javax.sound.sampled.AudioFormat;
import javax.sound.sampled.AudioSystem;
import javax.sound.sampled.DataLine;
import javax.sound.sampled.LineUnavailableException;
import javax.sound.sampled.SourceDataLine;
import javax.sound.sampled.TargetDataLine;

public class VoiceChat {

	private TargetDataLine microphone;
	private SourceDataLine speakers;
	private AudioFormat format;
	private byte[] micData;
	
	private int dataWidth;

	public VoiceChat() {
		format = new AudioFormat(8000.0f, 16, 1, true, true);

		try {

			// init microphone
			DataLine.Info info = new DataLine.Info(TargetDataLine.class, format);
			microphone = (TargetDataLine) AudioSystem.getLine(info);
			micData = new byte[microphone.getBufferSize() / 5];

			dataWidth = micData.length;
			
			// init speakers
			DataLine.Info dataLineInfo = new DataLine.Info(SourceDataLine.class, format);
			speakers = (SourceDataLine) AudioSystem.getLine(dataLineInfo);

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
	
	/**
	 * a test method that will play the recording back to itself
	 */
	public void playbackSelf(){
		int numBytesRead = 0;
		int bytesRead = 0;
		byte[] data = new byte[microphone.getBufferSize() / 5];
		int CHUNK_SIZE = 1024;
		
		while (bytesRead < 100000) {
			
			
			numBytesRead = microphone.read(data, 0, CHUNK_SIZE);
			bytesRead += numBytesRead;
			// write the mic data to a stream for use later
			// write mic data to stream for immediate playback
			speakers.write(data, 0, numBytesRead);
		}
	}
	
}
