package com.jangle.client;

import java.util.HashMap;

/**
 * Created by jess on 11/26/2016.
 */
public class Server {

    private int id;
    private HashMap<Integer, Channel> channels;
    private String avatarURL;
    private String name;

    public Server(int id){
        this.id = id;
        channels = new HashMap<>();
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public HashMap<Integer, Channel> getChannels() {
        return channels;
    }

    public String getAvatarURL() {
        return avatarURL;
    }

    public void setAvatarURL(String avatarURL) {
        this.avatarURL = avatarURL;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void addChannel(Channel channel) {
        if (channels.get(channel.getId()) == null)
            channels.put(channel.getId(), channel);
    }

    public Channel getChannel(int id) {
        return channels.get(id);
    }


}
