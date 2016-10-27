package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.client.User;
import com.jangle.communicate.Client_ParseData;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Button;
import javafx.scene.control.ListView;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;

import javafx.event.ActionEvent;
import javafx.stage.FileChooser;

import java.io.File;
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
    //TODO: Phase out userThread
    //private userThread userThread;
    private ObservableList<Message> testlist;

    @FXML
    public ListView<Message> messageArea;
    @FXML
    private TextField messageStage;
    @FXML
    private ListView<User> users;
    @FXML
    private Button attachButton;

    @FXML
    private void handleSendMessage(ActionEvent actionEvent) {
        String message = messageStage.getText();
        if (message.equals("Gimmie dat shit")){
            try {
                mClientParseData.request50MessagesWithOffset(0);
                messageStage.clear();
                return;
            } catch (IOException e) {
                e.printStackTrace();
                messageStage.clear();
                return;
            }
        }
        // Send the string to the server
        try {
            mClientParseData.sendMessage(new Message(1, message, 1, 1));
        } catch (IOException e) {
            e.printStackTrace();
        }
        messageStage.clear();
    }

    @FXML
    private void handleAttachment(ActionEvent actionEvent) {

        FileChooser fileChooser = new FileChooser();
        fileChooser.setTitle("Choose a file to attach.");
        File attachment = fileChooser.showOpenDialog(messageArea.getScene().getWindow());

        System.out.println(attachment);
        String[] splitPath = attachment.getAbsolutePath().split(".");
        if (splitPath.length != 2){
            //more than one period in file path
            //TODO: open an alert dialog
        }
        else {
            String extension = splitPath[1];
            if (extension.equals("png") || extension.equals("jpeg") || extension.equals("jpg") || extension.equals("bmp") || extension.equals("gif")) {
                //TODO: upload the file to the hosting site
            }
            else {
                //TODO: open a alert dialog
            }
        }
    }

    @Override
    public void initialize(URL location, ResourceBundle resources) {
        testlist = FXCollections.observableArrayList();

        //TODO: Phase out the userThread
        //userThread = new userThread(mClient, this);
    }

    public void updateMessages(ObservableList messages) {
        this.messageArea.setItems(messages);
    }

    public void updateUsers(ObservableList userList){
        this.users.setItems(userList);
    }

    public void setmClientParseData(Client_ParseData clientParseData){
        this.mClientParseData = clientParseData;
        this.mClient = mClientParseData.getClient();
        this.messageThread = new messageThread(mClient, this);
    }
}
