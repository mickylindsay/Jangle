package com.jangle.UI;

import com.cloudinary.Cloudinary;
import com.cloudinary.utils.ObjectUtils;
import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.communicate.Client_ParseData;
import javafx.event.ActionEvent;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Alert;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.stage.FileChooser;

import javax.swing.*;
import java.io.File;
import java.io.IOException;
import java.net.URL;
import java.util.Map;
import java.util.ResourceBundle;

/**
 * Created by sable_000 on 10/31/2016.
 */
public class settingsController implements Initializable {

    private ConfigUtil mConfigUtil;
    private Client mClient;
    private Client_ParseData mClientParseData;

    private ImageView backgroundImageView;
    @FXML
    private Button backgroundButton;
    @FXML
    private Label pathToBackground;
    @FXML
    private ImageView userIcon;
    @FXML
    private Button userIconButton;


    @Override
    public void initialize(URL location, ResourceBundle resources) {

    }

    @FXML
    private void handleBackgroundChange(ActionEvent actionEvent) {
        FileChooser fileChooser = new FileChooser();
        fileChooser.setTitle("Choose a new Image to use as a background");
        File newBackground = fileChooser.showOpenDialog(backgroundButton.getScene().getWindow());
        backgroundImageView.setImage(new Image(newBackground.toURI().toString()));
        mConfigUtil.setBackgroundPath(newBackground.getAbsolutePath());
        pathToBackground.setText(newBackground.getAbsolutePath());
    }

    @FXML
    private void handleUserIconChange(ActionEvent actionEvent) {
        FileChooser fileChooser = new FileChooser();
        fileChooser.setTitle("Choose a new Image to use as a Profile picture 32x32 res recommended");
        File newIcon = fileChooser.showOpenDialog(backgroundButton.getScene().getWindow());

        if (newIcon == null)
            return;

        String[] splitPath = newIcon.getAbsolutePath().split("\\.");

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
                    Map uploadResult = cloudinary.uploader().upload(newIcon, ObjectUtils.emptyMap());
                    String uploadURL = (String) uploadResult.get("url");
                    mClientParseData.sendNewUserIcon(uploadURL);
                } catch (IOException e) {
                    e.printStackTrace();
                }


            }
            else {
                createAttatchmentFileFormatErrorDialog();
            }
        }
    }

    public void setConfigUtil(ConfigUtil configUtil){
        this.mConfigUtil = configUtil;
        pathToBackground.setText(mConfigUtil.getBackgroundPath());
    }

    public void setmClient(Client client) {
        this.mClient = client;
        userIcon.setImage(new Image(mClient.findUser(mClient.getUserID()).getAvatarURL()));
    }

    public void setClientParseData(Client_ParseData clientParseData) {
        mClientParseData = clientParseData;
    }

    public void setBackgroundImageView(ImageView backgroundImageView) {
        this.backgroundImageView = backgroundImageView;
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


    public void handleCreateChannel(ActionEvent actionEvent) {
        String newName = JOptionPane.showInputDialog("Please enter the new name for the Channel");
        if(newName.length() > 20) {
            showServerNameTooLongAlert();
            return;
        }
        mClientParseData.createNewChannel(newName);
    }

    private void showServerNameTooLongAlert() {
        Alert alert = new Alert(Alert.AlertType.ERROR);
        alert.setTitle("Name too long!!");
        alert.setContentText("Channel name exceeded 20 Characters!");
        alert.showAndWait();
    }

    public void handleNameChange(ActionEvent actionEvent) {
        String newName = JOptionPane.showInputDialog("Please enter the new name for yourself");
        if (newName.length() > 20) {
            showDisplayNameTooLongAlert();
            return;
        }

        try {
            mClientParseData.setNewDisplayNameOnServer(newName);
        } catch (IOException e) {
            e.printStackTrace();
        }

    }

    private void showDisplayNameTooLongAlert() {
        Alert alert = new Alert(Alert.AlertType.ERROR);
        alert.setTitle("Name too long!!");
        alert.setContentText("Display name exceeded 20 Characters!");
        alert.showAndWait();
    }

}
