package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.communicate.Client_ParseData;

/**
 * Created by Jess on 10/13/2016.
 */
public class loginThread implements Runnable{

    private Client_ParseData mClient;
    private loginController mLoginController;


    public loginThread(Client_ParseData client, loginController loginController){
        this.mClient = client;
        this.mLoginController = loginController;

        Thread t = new Thread(this);
        t.start();
    }


    @Override
    public void run() {
       while (true) {
            if (mClient.getClient().getUserID() != 0) {
                mLoginController.successfulLogin();
                return;
            }
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
