package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.communicate.CommUtil;
import com.jangle.communicate.CommUtil.*;
import javafx.application.Platform;
import javafx.collections.FXCollections;

/**
 * Created by Jess on 11/16/2016.
 */
public class loginThread implements Runnable {

    private Client mClient;
    private loginController mLoginController;
    private Thread t;

    public loginThread(loginController controller, Client client) {
        this.mLoginController = controller;
        this.mClient = client;
        this.t = new Thread(this);
        t.start();
    }

    @Override
    public void run() {

        while (mClient.getLoginResult() != LoginResult.SUCESS && mClient.getUserID() == 0) {
            if (mClient.getLoginResult() == LoginResult.TIMEOUT && System.currentTimeMillis() - mClient.getLoginTime() > 3000){
                Platform.runLater(new Runnable() {
                    @Override
                    public void run() {
                        mLoginController.clearScreen();
                        mLoginController.loginTimeout.setVisible(true);
                    }
                });
            }
            else if (mClient.getLoginResult() == LoginResult.FAIL) {
                Platform.runLater(new Runnable() {
                    @Override
                    public void run() {
                        mLoginController.clearScreen();
                        mLoginController.failedLogin.setVisible(true);
                    }
                });
            }
            else if (mClient.getLoginResult() == LoginResult.NAME_TAKEN){
                Platform.runLater(new Runnable() {
                    @Override
                    public void run() {
                        mLoginController.clearScreen();
                        mLoginController.usernameTaken.setVisible(true);
                    }
                });
            }
            else if (mClient.getLoginResult() == LoginResult.SUCESS){
                System.out.print("here");
                if (mClient.getUserID() != 0) {
                    break;
                }
            }
            try {
                Thread.sleep(200);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

        }
        Platform.runLater(new Runnable() {
            @Override
            public void run() {
                mLoginController.successfulLogin();
            }
        });
    }

    public void stopThread() {
        t.interrupt();
    }
}
