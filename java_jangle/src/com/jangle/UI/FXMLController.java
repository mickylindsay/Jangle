package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.communicate.Client_ParseData;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.ListView;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;

import javafx.event.ActionEvent;
import java.io.IOException;
import java.net.URL;
import java.util.ResourceBundle;

/**
 * Created by Jess on 10/3/2016.
 */
public class FXMLController implements Initializable {

    private Client_ParseData mClientParseData;
    private Client mClient;
    private messageThread messageThread;
    private userThread userThread;

    @FXML
    public TextArea messageArea;
    @FXML
    private TextField messageStage;
    @FXML
    private ListView users;

    @FXML
    private void handleSendMessage(ActionEvent actionEvent) {
        String message = messageStage.getText();
        // Send the string to the server
        try {
            mClientParseData.sendMessage(new Message(0, message, System.currentTimeMillis(), 0, 0));
        } catch (IOException e) {
            e.printStackTrace();
        }
        messageArea.appendText(message);
        messageStage.clear();
    }

    @Override
    public void initialize(URL location, ResourceBundle resources) {

        mClient = new Client();

        try {
            mClientParseData = new Client_ParseData(mClient, "localhost", 9090);
        } catch (IOException e) {
            e.printStackTrace();
        }

        messageThread = new messageThread(mClient, this);
        userThread = new userThread(mClient, this);

    }

    public void addMessage(Message message) {
        messageArea.appendText(message.getMessageContent() + " " + message.getTimeStamp());
    }
}
