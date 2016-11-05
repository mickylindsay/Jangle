package com.jangle.UI;

import com.cloudinary.Cloudinary;
import com.cloudinary.utils.ObjectUtils;
import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.client.User;
import com.jangle.communicate.Client_ParseData;
import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.event.EventHandler;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.geometry.Pos;
import javafx.scene.control.*;

import javafx.event.ActionEvent;
import javafx.scene.control.Button;
import javafx.scene.control.TextField;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.scene.input.MouseEvent;
import javafx.scene.text.TextAlignment;
import javafx.stage.FileChooser;

import javax.print.DocFlavor;
import java.awt.*;
import java.io.File;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URL;
import java.util.Map;
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
            mClientParseData.sendMessage(new Message(mClient.getUserID(), message, 1, 1));
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
        if (attachment == null)
            return;

        String[] splitPath = attachment.getAbsolutePath().split("\\.");
        System.out.print(splitPath.length);
        //for (int i = 0; i<splitPath.length; i++)
            //System.out.println(splitPath[i]);

        if (splitPath.length != 2){
            //more than one period in file path
            Alert alert = new Alert(Alert.AlertType.ERROR);
            alert.setTitle("Invalid File Path");
            alert.setHeaderText("You chose an invalid file path");
            alert.setContentText("Error: (> 1 . in file path) Make sure that none of the folders are hidden.");
            alert.showAndWait();
        }
        else {
            String extension = splitPath[1];
            if (extension.equals("png") || extension.equals("jpeg") || extension.equals("jpg") || extension.equals("bmp") || extension.equals("gif")) {
                //TODO: upload the file to the hosting site
                //Cloudinary maven path: cloudinary-http
                Cloudinary cloudinary = new Cloudinary(ObjectUtils.asMap("cloud_name", "jangle", "api_key", "786816698113964", "api_secret", "vFTEtCmW_tOWLyXAia19UtIude4"));
                try {
                    Map uploadResult = cloudinary.uploader().upload(attachment, ObjectUtils.emptyMap());
                    String uploadURL = (String) uploadResult.get("url");
                    System.out.print(uploadURL);
                    mClientParseData.sendMessage(new Message(mClient.getUserID(), uploadURL, 1, 1));
                } catch (IOException e) {
                    e.printStackTrace();
                }


            }
            else {
                Alert alert = new Alert(Alert.AlertType.ERROR);
                alert.setTitle("Invalid Filetype");
                alert.setHeaderText("You chose an filetype that is not yet supported.");
                alert.setContentText("Error: The only currently supported filetpyes are: png, jpeg, jpg, bmp and gif");
                alert.showAndWait();
            }
        }
    }

    @Override
    public void initialize(URL location, ResourceBundle resources) {
        testlist = FXCollections.observableArrayList();

        messageArea.setCellFactory(listView -> new ListCell<Message>() {
            private ImageView imageView = new ImageView();
            @Override
            public void updateItem(Message message, boolean empty) {
                super.updateItem(message, empty);
                if (empty) {
                    setText(null);
                    setGraphic(null);
                } else {
                    if (message.isImg()) {
                        Image image = new Image(message.getMessageContent());
                        imageView.setImage(image);
                        imageView.setPreserveRatio(true);
                        imageView.setFitWidth(500);
                        setGraphic(imageView);
                        setContentDisplay(ContentDisplay.BOTTOM);
                        setAlignment(Pos.CENTER_LEFT);
                        //setTextAlignment(TextAlignment.LEFT);
                    }
                    setText(message.toString());
                }
            }
        });
        initializeListViewEventHandler();
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

    private void initializeListViewEventHandler(){
        messageArea.setOnMouseClicked(new EventHandler<MouseEvent>() {
            @Override
            public void handle(MouseEvent event) {
                if (messageArea.getSelectionModel().getSelectedItem().isImg() && Desktop.isDesktopSupported()){
                    try {
                        Desktop.getDesktop().browse(new URI(messageArea.getSelectionModel().getSelectedItem().getMessageContent()));
                    } catch (IOException e) {
                        e.printStackTrace();
                    } catch (URISyntaxException e) {
                        e.printStackTrace();
                    }
                }
            }
        });
    }
}
