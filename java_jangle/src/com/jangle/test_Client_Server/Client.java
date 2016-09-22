package com.jangle.test_Client_Server;

import com.jangle.communicate.network_Connection;

import java.io.Serializable;
import java.util.function.Consumer;

/**
 * Created by Jess on 9/22/2016.
 */

public class Client extends network_Connection {

    private String ip;
    private int port;

    public Client(String ip, int port, Consumer<Serializable> onRecieveCallback) {
        super(onRecieveCallback);
        this.ip = ip;
        this.port = port;
    }

    @Override
    protected boolean isServer() {
        return false;
    }

    @Override
    protected String getIP() {
        return ip;
    }

    @Override
    protected int getPort() {
        return port;
    }
}
