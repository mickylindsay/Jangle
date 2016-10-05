<<<<<<< HEAD
package com.jangle.UI;

import com.jangle.client.Client;
import com.sun.xml.internal.ws.policy.privateutil.PolicyUtils;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.Initializable;
import javafx.scene.control.TextArea;

/**
 * Created by sable_000 on 9/29/2016.
 */
public class messageThread implements Runnable {

    private Client mClient;
    private FXMLController ui;
    private ObservableList<String> messages;


    public messageThread(Client client, FXMLController ui){
        this.mClient = client;
        this.ui = ui;
        this.messages = FXCollections.observableArrayList();

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
                    messages.add(mClient.getMessages().get(mClient.getMessages().size() - difference + i).getMessageContent());
                    ui.updateMessages(messages);
                }
                size = mClient.getMessages().size();
            }

        }
    }
}
=======
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
            }

            else if (size < mClient.getMessages().size()){
                int difference = mClient.getMessages().size() - size;
                for (int i = 0; i < difference; i++) {
                    String message = mClient.getMessages().get(mClient.getMessages().size() - difference + i).getMessageContent();
                    ui.chatArea.appendText("Server: " + message + "\n");
                }
                size = mClient.getMessages().size();
            }

        }
    }
}
>>>>>>> 65e34f6c4b9512c07f461e6dd8c2b023093ae2ec
