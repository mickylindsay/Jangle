package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.client.User;
import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;
import javafx.application.Platform;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.scene.control.TextArea;

import java.util.ArrayList;

/**
 * Created by sable_000 on 9/29/2016.
 */
public class userThread implements Runnable {

    private Client mClient;
    private FXMLController ui;
    private ArrayList<User> mUsers;
    private ObservableList<User> mUserList;

    public userThread(Client client, FXMLController ui){
        this.mClient = client;
        this.ui = ui;
        this.mUserList = FXCollections.observableArrayList();
        this.mUsers = new ArrayList<>();

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
                    mUsers.add(mClient.getUsers().get(mClient.getUsers().size() - difference + i));
                    //TODO: Add user to user list here
                    //TODO: Display name updates and caching
                    Platform.runLater(new Runnable() {
                        @Override
                        public void run() {
                            mUserList = FXCollections.observableArrayList(mUsers);
                            ui.updateUsers(mUserList);
                        }
                    });
                }
                size = mClient.getUsers().size();
            }

        }
    }
}
