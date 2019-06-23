package gg.strims.chat.client;

import android.util.Log;
import androidx.appcompat.app.AppCompatActivity;
import org.java_websocket.client.WebSocketClient;
import org.java_websocket.handshake.ServerHandshake;

import android.os.Bundle;

import java.net.URI;
import java.net.URISyntaxException;

public class MainActivity extends AppCompatActivity {
    private WebSocketClient client;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
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
                Log.i("websocket",s);
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
