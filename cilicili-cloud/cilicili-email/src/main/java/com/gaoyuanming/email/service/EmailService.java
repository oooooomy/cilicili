package com.gaoyuanming.email.service;

import com.gaoyuanming.email.model.Message;

public interface EmailService {

    void send(Message message);

    void sendCode(String to) throws Exception;

    void checkCode(String email, String code) throws Exception;

}
