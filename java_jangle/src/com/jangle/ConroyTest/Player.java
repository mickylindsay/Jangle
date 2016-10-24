package com.jangle.ConroyTest;


import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import javax.sound.sampled.SourceDataLine;

public class Player extends Thread{
    public DatagramSocket din;
    public SourceDataLine audio_out;
    byte[] buffer = new byte[1024];
    @Override
    public void run(){
        DatagramPacket incoming  = new DatagramPacket(buffer, buffer.length);
        while (Server.calling) {            
            try {
                din.receive(incoming);
                buffer = incoming.getData();
                audio_out.write(buffer, 0, buffer.length);
            } catch (IOException ex) {
                System.out.Println(ex);
            }
        }
        audio_out.close();
        audio_out.drain();
    }
}