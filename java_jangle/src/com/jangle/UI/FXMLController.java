package com.jangle.UI;

import com.cloudinary.Cloudinary;
import com.cloudinary.utils.ObjectUtils;
import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.client.Server;
import com.jangle.client.User;
import com.jangle.communicate.Client_ParseData;
import com.jangle.voice.VoiceChat;

import javafx.collections.FXCollections;
import javafx.collections.ObservableList;
import javafx.event.EventHandler;
import javafx.fxml.FXML;
import javafx.fxml.FXMLLoader;
import javafx.fxml.Initializable;
import javafx.geometry.Pos;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.*;

import javafx.event.ActionEvent;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.TextField;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.scene.input.MouseButton;
import javafx.scene.input.MouseEvent;
import javafx.scene.layout.AnchorPane;
import javafx.scene.web.WebView;
import javafx.stage.FileChooser;
import javafx.stage.Stage;

import javax.swing.*;
import java.awt.*;
import java.io.File;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URL;
import java.util.*;

/**
 * Created by Jess on 10/3/2016.
 */
public class FXMLController implements Initializable {

    private Client_ParseData mClientParseData;
    private Client mClient;
    private messageThread mMessageThread;
    private ConfigUtil mConfigUtil;
    private ObservableList<Message> testlist;
    private VoiceChat mVoice;

    @FXML
    public ListView<Message> messageArea;
    @FXML
    private TextField messageStage;
    @FXML
    private ListView<User> userList;
    @FXML
    private Button attachButton;
    @FXML
    private Button settingsButton;
    @FXML
    protected ImageView chatBackground;
    @FXML
    private ListView<Server> serverList;
    @FXML
    private Button connectButton;
    @FXML
    private Button muteButton;
    @FXML
    private Label loadingLabel;
    @FXML
    private AnchorPane loadingPane;
    @FXML
    private ImageView loadingImage;


    @FXML
    private void handleSendMessage(ActionEvent actionEvent) {
        //TODO: Carriage return algorithm
        String message = messageStage.getText();

        if (message.length() == 0)
            return;
        // Send the string to the server
        try {
            mClientParseData.sendMessage(new Message(mClient.getUserID(), message, mClient.getCurrentServerID(), mClient.getCurrentChannelID()));
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

        //System.out.println(attachment);
        if (attachment == null)
            return;

        String[] splitPath = attachment.getAbsolutePath().split("\\.");
        System.out.print(splitPath.length);
        //for (int i = 0; i<splitPath.length; i++)
            //System.out.println(splitPath[i]);

        if (splitPath.length != 2){
            createAttatchmentErrorDialog();
        }
        else {
            String extension = splitPath[1];
            extension = extension.toLowerCase();
            if (extension.equals("png") || extension.equals("jpeg") || extension.equals("jpg") || extension.equals("bmp") || extension.equals("gif")) {
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
                createAttatchmentFileFormatErrorDialog();
            }
        }
    }


    @FXML
    private void handleSettings(ActionEvent actionEvent) {
        Stage settingsStage = new Stage();
        settingsStage.setScene(new Scene(createSettingsDialog()));
        settingsStage.showAndWait();
        refresh();
    }



    @Override
    public void initialize(URL location, ResourceBundle resources) {
        testlist = FXCollections.observableArrayList();

        setMessageAreaCellFactory();
        setServerListCellFactory();
        setUserListCellFactory();


        initializeListViewEventHandler();

    }

    private void setServerListCellFactory() {
        serverList.setCellFactory(listView -> new ListCell<Server>() {
            private ImageView imageview = new ImageView();
            @Override
            public void updateItem(Server server, boolean empty) {
                super.updateItem(server, empty);
                if(empty){
                    setText(null);
                    setGraphic(null);
                    return;
                }
                else {

                }
            }
        });
    }

    private void setUserListCellFactory() {
        {
            userList.setCellFactory(listView -> new ListCell<User>() {
                private ImageView imageView = new ImageView();
                @Override
                public void updateItem(User user, boolean empty) {
                    super.updateItem(user, empty);
                    if (empty) {
                        setText(null);
                        setGraphic(null);
                        return;
                    }
                    else if(user.getDisplayName() == null) {
                        try {
                            mClientParseData.requestDisplayName(user);
                            mClientParseData.requestAvatarURL(user);
                        } catch (IOException e) {
                            e.printStackTrace();
                        }
                    }
                    else {
                        if (user.isChannel()) {
                            setGraphic(null);
                            setText(user.toString());
                        } else {
                            Image image;
                            if (user.getChannelID() == 0) {
                                image = new Image(user.OFFLINE_AVATAR, 20, 20, false, true);
                            }
                            else {
                                if (isImg(user.getAvatarURL()))
                                    image = user.getImage();
                                else
                                    image = new Image(user.DEFAULT_AVATAR, 20, 20, false, true);
                            }
                            imageView.setImage(image);
                            setGraphic(imageView);
                            setContentDisplay(ContentDisplay.LEFT);
                            setAlignment(Pos.CENTER_LEFT);
                            //setTextAlignment(TextAlignment.LEFT);
                        }
                    }
                    if(user.isChannel())
                        setText(user.getChannel().toString());
                    else
                        setText(user.getDisplayName());
                }
            });
        }
    }

    private void setMessageAreaCellFactory() {
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
                        imageView.setImage(message.getImage());
                        setGraphic(imageView);
                        setContentDisplay(ContentDisplay.BOTTOM);
                        setAlignment(Pos.CENTER_LEFT);
                        //setTextAlignment(TextAlignment.LEFT);
                    }
                    else if (message.isYoutube() && message.isPlaying()){
                        //TODO: add click to play
                        setGraphic(message.getWebView());
                        setContentDisplay(ContentDisplay.BOTTOM);
                        setAlignment(Pos.CENTER_LEFT);
                    }

                    else if (message.isYoutube() && !message.isPlaying()){
                        setGraphic(new ImageView(message.getImage()));
                        setContentDisplay(ContentDisplay.BOTTOM);
                        setAlignment(Pos.CENTER_LEFT);
                    }
                    else
                        setGraphic(null);
                    setText(formatMessage(message));

                }
            }
        });
    }

    public void updateMessages(ObservableList messages) {
        this.messageArea.setItems(messages);
    }

    public void updateUsers(ObservableList userList){
        this.userList.setItems(userList);
    }

    public void setmClientParseData(Client_ParseData clientParseData){
        this.mClientParseData = clientParseData;
        this.mClient = mClientParseData.getClient();
        this.mMessageThread = new messageThread(mClient, this);
    }

    public void setConfigUtil(ConfigUtil configUtil){
        this.mConfigUtil = configUtil;
        if (mConfigUtil.getBackgroundPath() != null)
            chatBackground.setImage(new Image(new File(mConfigUtil.getBackgroundPath()).toURI().toString()));
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
                else if(messageArea.getSelectionModel().getSelectedItem().isYoutube() && Desktop.isDesktopSupported() && messageArea.getSelectionModel().getSelectedItem().isPlaying()){
                    try {
                        Desktop.getDesktop().browse(new URI(messageArea.getSelectionModel().getSelectedItem().getMessageContent()));
                    } catch (IOException e) {
                        e.printStackTrace();
                    } catch (URISyntaxException e) {
                        e.printStackTrace();
                    }
                }
                else if (messageArea.getSelectionModel().getSelectedItem().isYoutube() && !messageArea.getSelectionModel().getSelectedItem().isPlaying()){
                    messageArea.getSelectionModel().getSelectedItem().setPlaying(true);
                    messageArea.refresh();
                }
            }
        });

        userList.setOnMouseClicked(new EventHandler<MouseEvent>() {
            @Override
            public void handle(MouseEvent event) {
                if (userList.getSelectionModel().getSelectedItem().isChannel()){
                    if(event.getButton() == MouseButton.SECONDARY){
                        changeChannelNameAlert(userList.getSelectionModel().getSelectedItem());
                        return;
                    }
                    for (Message m: mClient.getMessages()) {
                        m.setPlaying(false);
                    }
                    messageArea.refresh();
                    mClient.changeChannel(userList.getSelectionModel().getSelectedItem().getId()-1000);
                    mClientParseData.changeLocation();
                    if (mClient.getMessages(mClient.getCurrentServerID(), mClient.getCurrentChannelID()).size() == 0) {
                        try {
                            mClientParseData.request50MessagesWithOffset(0);
                        } catch (IOException e) {
                            e.printStackTrace();
                        }
                    }
                    //Change the messages to the ones in the current channel
                    updateMessages(FXCollections.observableArrayList(mClient.getMessages(mClient.getCurrentServerID(), mClient.getCurrentChannelID())));
                    updateUsers(FXCollections.observableList(mClient.getUsers()));
                }
            }
        });

    }



    private Parent createSettingsDialog() {
        FXMLLoader loader = new FXMLLoader(getClass().getResource("res/fxml/settings.fxml"));
        AnchorPane dialog = null;
        try {
            dialog = loader.load();
        } catch (IOException e) {
            e.printStackTrace();
        }
        settingsController mSettings = loader.getController();
        mSettings.setConfigUtil(mConfigUtil);
        mSettings.setmClient(mClient);
        mSettings.setClientParseData(mClientParseData);
        mSettings.setBackgroundImageView(chatBackground);

        return dialog;
    }

    private String formatMessage(Message message) {
        User sender = mClient.findUser(message.getUserID());

        if(sender == null)
            return message.getUserID() + "          " + message.getTime() +"\n" + message.getMessageContent();

        return sender.getDisplayName() + "          " + message.getTime() + "\n" + message.getMessageContent();
    }

    @FXML
    public void handleMute(ActionEvent actionEvent) {
    	if (!mClient.isConnectedToVoice()){
    		return;
    	}
    	
    	if (mClient.getBroadcastStatus()){
            mVoice.endBrodcast();
            muteButton.setText("Unmute");
    		
    	}
    	else{
    		mVoice.startBrodcast();
            muteButton.setText("Mute");
    	}
    }

    @FXML
    public void handleVoipConnection(ActionEvent actionEvent) {
    	
    	if (mClient.isConnectedToVoice()){
    		mVoice.disconnectFromVoice();
            connectButton.setText("Connect");
    	}
    	else{
    		mVoice.connectToVoice();
            connectButton.setText("Disconnect");
    	}
    }
    
    public void setVoiceChat(VoiceChat gVoice){
    	this.mVoice = gVoice;
    }

    private void createAttatchmentErrorDialog() {
        Alert alert = new Alert(Alert.AlertType.ERROR);
        alert.setTitle("Invalid File Path");
        alert.setHeaderText("You chose an invalid file path");
        alert.setContentText("Error: (> 1 . in file path) Make sure that none of the folders are hidden.");
        alert.showAndWait();
    }

    private void createAttatchmentFileFormatErrorDialog() {
        Alert alert = new Alert(Alert.AlertType.ERROR);
        alert.setTitle("Invalid Filetype");
        alert.setHeaderText("You chose an filetype that is not yet supported.");
        alert.setContentText("Error: The only currently supported filetpyes are: png, jpeg, jpg, bmp and gif");
        alert.showAndWait();
    }

    private void changeChannelNameAlert(User user) {
        String newName = JOptionPane.showInputDialog("Please enter a new name for the channel:");
        if (newName == null)
            return;
        else if (newName.length() > 20){
            openTooLongAlert();
            return;
        }

        mClientParseData.sendNewChannelName(user.getChannelID(), newName);
    }

    private void openTooLongAlert() {
        Alert alert = new Alert(Alert.AlertType.ERROR);
        alert.setTitle("That name is too long!");
        alert.setContentText("Yo dawg yo name is more than 20 characters. Make dat shorter");
        alert.showAndWait();
    }

    public boolean isImg(String s) {
        return s.contains("http://") && (s.contains(".png") || s.contains(".jpg") || s.contains(".gif") || s.contains("jpeg") || s.contains(".bmp"));
    }

    public messageThread getMessageThread() {
        return mMessageThread;
    }

    public void finishedLoading() {
        messageArea.refresh();
        userList.refresh();
        loadingImage.setVisible(false);
        loadingLabel.setVisible(false);
        loadingPane.setVisible(false);
    }

    public void refresh() {
        messageArea.refresh();
        userList.refresh();
    }
}
