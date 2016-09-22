package com.jangle.test_Client_Server;

import com.jangle.communicate.network_Connection;

import java.io.Serializable;
import java.util.function.Consumer;

/**
 * Created by Jess on 9/22/2016.
 */

public class Server extends network_Connection {

    private int port;

    public Server(int port, Consumer<Serializable> onRecieveCallback) {
        super(onRecieveCallback);
        this.port = port;
    }

    @Override
    protected boolean isServer() {
        return true;
    }

    @Override
    protected String getIP() {
        return null;
    }

    @Override
    protected int getPort() {
        return port;
    }
}
