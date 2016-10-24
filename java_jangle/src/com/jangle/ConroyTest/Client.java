package com.jangle.ConroyTest;


import com.sun.security.ntlm.Client;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;
import java.net.UnknownHostException;
import javax.sound.sampled.AudioFormat;
import javax.sound.sampled.AudioSystem;
import javax.sound.sampled.DataLine;
import javax.sound.sampled.LineUnavailableException;
import javax.sound.sampled.TargetDataLine;

public class Client {
		public static boolean calling = false;

    public int port_server = 9090;
    public String add_server = "localhost";
		TargetDataLine audio_in;
		
    public static AudioFormat getaudioformat(){
        float sampleRate = 8000.0F;
        int sampleSizeInbits = 16;
        int channel = 2;
        boolean signed = true;
        boolean bigEndian = false;
        return new AudioFormat(sampleRate, sampleSizeInbits, channel, signed, bigEndian);
    }
    
    public Client() {
			 try {
            AudioFormat format = getaudioformat();
            DataLine.Info info = new DataLine.Info(TargetDataLine.class, format);
            if(!AudioSystem.isLineSupported(info)){
                System.exit(0);
            }
            audio_in = (TargetDataLine) AudioSystem.getLine(info);
            audio_in.open(format);
            audio_in.start();
            Recorder r = new Recorder();
            InetAddress inet = InetAddress.getByName(add_server);
            r.audio_in = audio_in;
            r.dout = new DatagramSocket();
            r.server_ip = inet;
            r.server_port = port_server;
            Client.calling = true;
            r.start();
            btn_start.setEnabled(false);
            btn_stop.setEnabled(true);
        } catch (LineUnavailableException | UnknownHostException | SocketException ex) {
            System.out.Println(ex);
        }
    }
		   
    public static void main(String[] args) {
        Client c = new Client();
    }

}