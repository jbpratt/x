package gg.strims.chat.client;

import android.app.Activity;
import android.view.LayoutInflater;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.TextView;
import android.content.Context;
import android.view.View;
import java.util.ArrayList;
import java.util.List;

public class MessageAdapter extends BaseAdapter {
    List<Message> messages = new ArrayList<Message>();
    Context context;

    public MessageAdapter(Context context) {
        this.context = context;
    }

    public void add(Message message) {
        this.messages.add(message);
        notifyDataSetChanged();
    }

    @Override
    public int getCount() {
        return messages.size();
    }

    @Override
    public Object getItem(int i) {
        return messages.get(i);
    }

    @Override
    public long getItemId(int i) {
        return i;
    }

    @Override
    public View getView(int i, View convertView, ViewGroup viewGroup) {
        MessageViewHolder holder = new MessageViewHolder();
        LayoutInflater messageInflater = (LayoutInflater) context.getSystemService(Activity.LAYOUT_INFLATER_SERVICE);
        Message msg = messages.get(i);

        convertView = messageInflater.inflate(R.layout.inbound_message, null);

        holder.name = (TextView) convertView.findViewById(R.id.name);
        holder.body = (TextView) convertView.findViewById(R.id.message_body);
        convertView.setTag(holder);

        holder.name.setText(msg.getUser());
        holder.body.setText(msg.getText());

        return convertView;
    }
}

class MessageViewHolder {
    public TextView name;
    public TextView body;
}