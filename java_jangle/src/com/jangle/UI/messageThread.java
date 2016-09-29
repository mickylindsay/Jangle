package com.jangle.UI;

import com.jangle.client.Client;
import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;
import javafx.scene.control.TextArea;

/**
 * Created by sable_000 on 9/29/2016.
 */
public class messageThread implements Runnable {

    private Client mClient;
    private Text_UI ui;

    public messageThread(Client client, Text_UI ui){
        this.mClient = client;
        this.ui = ui;

        Thread t = new Thread(this);
        t.start();
    }

    @Override
    public void run() {
        int size = 0;
        while(true) {

            if (size == mClient.getMessages().size()){
                try {
                    Thread.sleep(200);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                continue;
            }
            else if (size < mClient.getMessages().size()){
                int difference = mClient.getMessages().size() - size;
                for (int i = 0; i < difference; i++) {
                    String message = mClient.getMessages().get(mClient.getMessages().size() - 1 - difference + i).getMessageContent();
                    ui.chatArea.appendText("Server: " + message + "\n");
                }
            }

        }
    }
}
