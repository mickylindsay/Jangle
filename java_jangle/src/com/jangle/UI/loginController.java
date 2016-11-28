package com.jangle.UI;

import com.jangle.client.Message;
import com.jangle.communicate.Client_ParseData;
import com.jangle.communicate.CommUtil;
import javafx.application.Platform;
import javafx.event.ActionEvent;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.PasswordField;
import javafx.scene.control.TextField;
import javafx.scene.image.ImageView;
import javafx.stage.Stage;

import java.io.IOException;
import java.net.URL;
import java.util.ResourceBundle;

/**
 * Created by Jess on 10/5/2016.
 */
public class loginController implements Initializable {

    private Client_ParseData mClient_parseData;
    //private loginThread mLoginThread;
    private CommUtil.LoginResult mLoginResult;
    private loginThread mLoginThread;

    @FXML
    public TextField usernameField;
    @FXML
    public PasswordField passwordField;
    @FXML
    public Button registerButton;
    @FXML
    public Button logInButton;
    @FXML
    public Label failedLogin;
    @FXML
    public Label loginTimeout;
    @FXML
    public Label itWontFitSenpai;
    @FXML
    public ImageView loadingAnim;
    @FXML
    public Label tooSmall;
    @FXML
    public Label noUsernameOrPassword;
    @FXML
    public Label usernameTaken;

    @FXML
    private void handleLogin(ActionEvent actionEvent) {
        String username = usernameField.getText();
        String password = passwordField.getText();
        clearScreen();

        if (username.length() > 20){
            itWontFitSenpai.setVisible(true);
            return;
        }
        else if(username.length() == 0 || password.length() == 0){
            noUsernameOrPassword.setVisible(true);
            return;
        }
        else if(username.length() < 3){
            tooSmall.setVisible(true);
            return;
        }

        //Send login to server
        loadingAnim.setVisible(true);

        try {
            mClient_parseData.submitLogIn(username, password);
        } catch (IOException e) {
            e.printStackTrace();
        }

    }

    @FXML
    private void handleRegister(ActionEvent actionEvent) {
        String username = usernameField.getText();
        String password = passwordField.getText();
        clearScreen();

        if (username.length() > 20){
            itWontFitSenpai.setVisible(true);
            return;
        }
        else if(username.length() == 0 || password.length() == 0){
            noUsernameOrPassword.setVisible(true);
            return;
        }
        else if(username.length() < 3){
            tooSmall.setVisible(true);
            return;
        }

        loadingAnim.setVisible(true
        );
        // Send the register user to the server

        try {
            mClient_parseData.createUserInServer(username, password);
        } catch (IOException e) {
            e.printStackTrace();
        }

    }

    @Override
    public void initialize(URL location, ResourceBundle resources) {

    }

    public void clearScreen(){
        loginTimeout.setVisible(false);
        itWontFitSenpai.setVisible(false);
        failedLogin.setVisible(false);
        loadingAnim.setVisible(false);
        tooSmall.setVisible(false);
        usernameTaken.setVisible(false);
        noUsernameOrPassword.setVisible(false);
    }

    public void successfulLogin() {
        Stage here = (Stage) logInButton.getScene().getWindow();
        here.close();
    }


    public void setmClient_parseData(Client_ParseData Client_parseData){
        this.mClient_parseData = Client_parseData;
        mLoginThread = new loginThread(this, mClient_parseData.getClient());
    }
}
