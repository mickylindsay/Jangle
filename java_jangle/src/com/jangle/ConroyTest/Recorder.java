package com.jangle.ConroyTest;


import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import javax.sound.sampled.TargetDataLine;


public class Recorder extends Thread{
	
    public TargetDataLine audio_in = null;
    public DatagramSocket dout;
    byte byte_buff[] = new byte[512];
    public InetAddress server_ip;
    public int server_port;
		
    @Override
    public void run(){
        while(Client_voice.calling){
            try {
                audio_in.read(byte_buff, 0, byte_buff.length);
                DatagramPacket data = new DatagramPacket(byte_buff, byte_buff.length, server_ip, server_port);
                dout.send(data);
            } catch (IOException ex) {
                System.out.Println(ex);
            }
        }
        audio_in.close();
        audio_in.drain();
		}
}
