package com.gaoyuanming.room.websocket;

import org.springframework.stereotype.Component;

import javax.websocket.OnClose;
import javax.websocket.OnError;
import javax.websocket.OnOpen;
import javax.websocket.server.PathParam;
import javax.websocket.server.ServerEndpoint;

@Component
@ServerEndpoint(value = "/chat/{roomId}/{userId}")
public class ChatWebSocket {

    @OnOpen
    public void onOpen(@PathParam("roomId") String roomId, @PathParam("userId") String userId) {
        System.out.println("onOpen roomId=" + roomId);
        System.out.println("onOpen userId=" + userId);
    }

    @OnClose
    public void onClose() {

    }

    @OnError
    public void onError(Throwable error) {
        System.out.println("onOpen roomId=");
    }

}
