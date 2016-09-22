package com.jangle.communicate;

import java.io.ObjectInput;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.Serializable;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.function.Consumer;

/**
 * Created by Jess on 9/22/2016.
 */
public abstract class network_Connection {

    private connectionThread mConnectionThread = new connectionThread();
    private Consumer<Serializable> onRecieveCallback;

    public network_Connection(Consumer<Serializable> onRecieveCallback){
        this.onRecieveCallback = onRecieveCallback;
        mConnectionThread.setDaemon(true);
    }

    public void startConnection() throws Exception {
        mConnectionThread.start();
    }

    public void send(Serializable data) throws Exception{
        mConnectionThread.out.writeObject(data);
    }

    public void closeConnection() throws Exception {
        mConnectionThread.mSocket.close();
    }

    protected abstract boolean isServer();
    protected abstract String getIP();
    protected abstract int getPort();

    private class connectionThread extends Thread {
        private Socket mSocket;
        private ObjectOutputStream out;

        @Override
        public void run() {
            try(ServerSocket server = isServer() ? new ServerSocket(getPort()) : null;
                Socket socket = isServer() ? server.accept() : new Socket(getIP(), getPort());
                ObjectOutputStream out = new ObjectOutputStream(socket.getOutputStream());
                ObjectInputStream in = new ObjectInputStream(socket.getInputStream())) {

                this.mSocket = socket;
                this.out = out;
                socket.setTcpNoDelay(true);

                while (true) {
                    Serializable data = (Serializable) in.readObject();
                    onRecieveCallback.accept(data);
                }
            }
            catch (Exception e) {
                System.out.println("Exception caught: " + e);
                onRecieveCallback.accept("Connection Closed. Error: " + e);
            }

        }
    }

}
