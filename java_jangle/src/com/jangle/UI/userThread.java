package com.jangle.UI;

import com.jangle.client.Client;
import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;
import javafx.scene.control.TextArea;

/**
 * Created by sable_000 on 9/29/2016.
 */
public class userThread implements Runnable {

    private Client mClient;
    private FXMLController ui;

    public userThread(Client client, FXMLController ui){
        this.mClient = client;
        this.ui = ui;

        Thread t = new Thread(this);
        t.start();
    }

    @Override
    public void run() {
        int size = 0;
        while(true) {

            if (size == mClient.getUsers().size()){
                try {
                    Thread.sleep(200);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

            else if (size < mClient.getUsers().size()){
                int difference = mClient.getUsers().size() - size;
                for (int i = 0; i < difference; i++) {
                    String message = mClient.getUsers().get(mClient.getUsers().size() - difference + i).getDisplayName();
                    //TODO: Add user to user list here
                    //ui.chatArea.appendText("Server: " + message + "\n");
                }
                size = mClient.getUsers().size();
            }

        }
    }
}
