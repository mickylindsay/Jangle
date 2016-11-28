package com.jangle.client;

import java.util.ArrayList;

/**
 * Created by jess on 11/26/2016.
 */
public class Channel {

    //Messages do not currently use a hash map because message Id is not in yet
    private ArrayList<Message> mMessages;
    private int id;
    private String name;

    public Channel(int id) {
        this.id = id;
        mMessages = new ArrayList<>();
        this.name = null;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public ArrayList<Message> getmMessages() {
        return mMessages;
    }

    public void addMessage(Message message) {
        if (mMessages.contains(message))
            return;
        //keeps messages chronologically ordered
        if (mMessages.size() == 0)
            mMessages.add(message);
        else {
            for (int i = 0; i < mMessages.size(); i++) {
                if (message.getTimeStamp() > mMessages.get(i).getTimeStamp())
                    continue;
                else if (message.getTimeStamp() == mMessages.get(i).getTimeStamp()) {
                    mMessages.add(i + 1, message);
                    return;
                }
                else {
                    mMessages.add(i, message);
                    return;
                }
            }
            //add at the end if the loop finishes
            mMessages.add(message);
        }
    }

    public String toString() {
        if (this.name == null)
            return "Channel: " + id;
        else
            return "Channel: " + name;
    }
}
