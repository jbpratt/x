package gg.strims.chat.client;

public class Message {

    private String text ;
    private String user ;
    private String type;

    public Message(String text, String user, String type) {
        this.text = text;
        this.user = user;
        this.type = type;
    }

    public String getText() {
        return text;
    }
    public String getUser() {
        return user;
    }
    public String getType() {
        return type;
    }
}
