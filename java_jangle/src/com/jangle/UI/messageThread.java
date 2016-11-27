package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.client.User;
import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;
import javafx.application.Platform;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.Initializable;
import javafx.scene.control.TextArea;

import java.util.ArrayList;

/**
 * Created by Jess on 9/29/2016.
 */
public class messageThread implements Runnable {

    private Client mClient;
    private FXMLController ui;
    private ObservableList<Message> messageList;
    private ArrayList<Message> messages;
    private ArrayList<User> mUsers;
    private ObservableList<User> mUserList;


    public messageThread(Client client, FXMLController ui){
        this.mClient = client;
        this.ui = ui;
        this.messageList = FXCollections.observableArrayList();
        this.messages = new ArrayList<>();
        this.mUserList = FXCollections.observableArrayList();
        this.mUsers = new ArrayList<>();
        Thread t = new Thread(this);
        t.start();
    }

    @Override
    public void run() {
        //Messages size
        int mSize = 0;
        //Users Size
        int uSize = 0;

        while(true) {
            System.out.println(mClient.getServer(1).getChannels().size());

            if (mSize == mClient.getMessages(mClient.getCurrentServerID(),mClient.getCurrentChannelID()).size()){
                try {
                    Thread.sleep(200);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

            else if (mSize < mClient.getMessages(mClient.getCurrentServerID(), mClient.getCurrentChannelID()).size()){
                int difference = mClient.getMessages(mClient.getCurrentServerID(), mClient.getCurrentChannelID()).size() - mSize;
                Message toDisplay = null;

                for (int i = 0; i < difference; i++) {
                    toDisplay = mClient.getMessages(mClient.getCurrentServerID(),mClient.getCurrentChannelID()).get(mClient.getMessages(mClient.getCurrentServerID(), mClient.getCurrentChannelID()).size() - difference + i);
                    messages.add(toDisplay);

                    //Making new UI update thread
                    Platform.runLater(new Runnable() {
                        @Override
                        public void run() {
                          	messageList = FXCollections.observableArrayList(mClient.getMessages(mClient.getCurrentServerID(), mClient.getCurrentChannelID()));
                            ui.updateMessages(messageList);
                        }
                    });
                }
                mSize = mClient.getMessages(1,1).size();
            }


            //Handling user listening
            if (uSize < mClient.getUsers().size()){
                int difference = mClient.getUsers().size() - uSize;
                for (int i = 0; i < difference; i++) {
                    mUsers.add(mClient.getUsers().get(mClient.getUsers().size() - difference + i));
                    //TODO: Display name updates and caching
                    Platform.runLater(new Runnable() {
                        @Override
                        public void run() {
                            mUserList = FXCollections.observableArrayList(mUsers);
                            ui.updateUsers(mUserList);
                        }
                    });
                }
                uSize = mClient.getUsers().size();
            }

        }
    }
}