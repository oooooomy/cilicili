package com.gaoyuanming.service.impl;

import com.gaoyuanming.model.entity.Account;
import com.gaoyuanming.service.AccountService;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import java.io.IOException;

@Service
public class AccountServiceImpl implements AccountService {

    @Value("${appId}")
    private String appId;

    @Value("${appSecret}")
    private String appSecret;

    @Override
    public Account login(String code) {
        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder()
                .url("https://api.weixin.qq.com/sns/jscode2session?appid=" + appId + "&secret=" + appSecret + "&js_code=" + code + "&grant_type=authorization_code")
                .build();
        Response response = null;
        try {
            response = client.newCall(request).execute();
            if (response.isSuccessful()) {
                System.out.println(response.body().string());
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        return null;
    }

}
