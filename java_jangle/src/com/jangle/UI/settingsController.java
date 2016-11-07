package com.jangle.UI;

import javafx.event.ActionEvent;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.stage.FileChooser;

import java.io.File;
import java.net.URL;
import java.util.ResourceBundle;

/**
 * Created by sable_000 on 10/31/2016.
 */
public class settingsController implements Initializable {

    private ConfigUtil mConfigUtil;
    private ImageView backgroundImageView;
    @FXML
    private Button backgroundButton;
    @FXML
    private Label pathToBackground;

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

    public void setConfigUtil(ConfigUtil configUtil){
        this.mConfigUtil = configUtil;
        pathToBackground.setText(mConfigUtil.getBackgroundPath());
    }

    public void setBackgroundImageView(ImageView backgroundImageView) {
        this.backgroundImageView = backgroundImageView;

    }


}
