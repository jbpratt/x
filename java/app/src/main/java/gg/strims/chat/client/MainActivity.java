package gg.strims.chat.client;

import android.util.Log;
import androidx.appcompat.app.AppCompatActivity;
import org.java_websocket.client.WebSocketClient;
import org.java_websocket.handshake.ServerHandshake;
import android.widget.ListView;
import android.os.Bundle;
import org.json.*;

import java.net.URI;
import java.net.URISyntaxException;

public class MainActivity extends AppCompatActivity {
    private WebSocketClient client;
    private MessageAdapter messageAdapter;
    private ListView messagesView;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        messageAdapter = new MessageAdapter(this);
        messagesView = (ListView) findViewById(R.id.messages_view);
        messagesView.setAdapter(messageAdapter);

        connect();
    }

    private void connect() {
        URI uri;
        try {
            uri = new URI("wss://chat.strims.gg/ws");
        } catch (URISyntaxException e) {
            e.printStackTrace();
            return;
        }

        client = new WebSocketClient(uri) {
            @Override
            public void onOpen(ServerHandshake handshake) {
                Log.i("websocket", "opened");
            }

            @Override
            public void onMessage(String s) {
                try {
                    JSONObject obj = new JSONObject(s);
                    //String user = obj.getJSONObject("msg").getString("user");
                    Log.i("websocket",obj.names().toString());
                } catch (JSONException e) {
                    e.printStackTrace();
                }

            }

            @Override
            public void onClose(int i, String s, boolean b) {
                Log.i("websocket", "closed");
            }

            @Override
            public void onError(Exception e) {
                Log.i("websocket", "Error " + e.getMessage());
            }
        };

        client.connect();

    }
}
