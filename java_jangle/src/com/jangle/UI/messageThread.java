package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.client.Message;
import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.Initializable;
import javafx.scene.control.TextArea;

import java.util.ArrayList;

/**
 * Created by sable_000 on 9/29/2016.
 */
public class messageThread implements Runnable {

    private Client mClient;
    private FXMLController ui;
    private ObservableList<Message> messageList;
    private ArrayList<Message> messages;


    public messageThread(Client client, FXMLController ui){
        this.mClient = client;
        this.ui = ui;
        this.messageList = FXCollections.observableArrayList();
        this.messages = new ArrayList<>();
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
            }

            else if (size < mClient.getMessages().size()){
                int difference = mClient.getMessages().size() - size;
                for (int i = 0; i < difference; i++) {
                    //String message = mClient.getMessages().get(mClient.getMessages().size() - difference + i).getMessageContent();
                    messages.add(mClient.getMessages().get(mClient.getMessages().size() - difference + i));
                    messageList = FXCollections.observableArrayList(messages);
                    ui.updateMessages(messageList);
                }
                size = mClient.getMessages().size();
            }

        }
    }
}