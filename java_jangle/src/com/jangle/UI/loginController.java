package com.jangle.UI;

import com.jangle.client.Message;
import com.jangle.communicate.Client_ParseData;
import javafx.event.ActionEvent;
import javafx.fxml.FXML;
import javafx.fxml.Initializable;
import javafx.scene.control.Button;
import javafx.scene.control.PasswordField;
import javafx.scene.control.TextField;

import java.io.IOException;
import java.net.URL;
import java.util.ResourceBundle;

/**
 * Created by Jess on 10/5/2016.
 */
public class loginController implements Initializable {

    private Client_ParseData mClient_parseData;

    @FXML
    public TextField usernameField;
    @FXML
    public PasswordField passwordField;
    @FXML
    public Button registerButton;
    @FXML
    public Button logInButton;

    @Override
    public void initialize(URL location, ResourceBundle resources) {

    }

    @FXML
    private void handleLogin(ActionEvent actionEvent) {
        String username = usernameField.getText();
        String password = passwordField.getText();

        //Send login to server
    }

    @FXML
    private void handleRegister(ActionEvent actionEvent) {
        String username = usernameField.getText();
        String password = passwordField.getText();
        // Send the register user to the server
    }
}
