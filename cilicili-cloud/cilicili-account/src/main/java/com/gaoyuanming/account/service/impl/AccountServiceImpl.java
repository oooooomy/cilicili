package com.gaoyuanming.account.service.impl;

import com.alibaba.fastjson.JSON;
import com.gaoyuanming.account.model.dto.LoginDto;
import com.gaoyuanming.account.model.entity.Account;
import com.gaoyuanming.account.repository.AccountRepository;
import com.gaoyuanming.account.service.AccountService;
import okhttp3.OkHttpClient;
import okhttp3.Request;
import okhttp3.Response;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.io.IOException;
import java.util.Objects;
import java.util.Optional;

@Service
public class AccountServiceImpl implements AccountService {

    @Value("${appId}")
    private String appId;

    @Value("${appSecret}")
    private String appSecret;

    @Resource
    private AccountRepository accountRepository;

    /**
     * 用户登录实现
     *
     * @param code 微信code
     * @return 用户信息
     * @throws IOException 异常
     */
    @Override
    public Account login(String code) throws IOException {
        OkHttpClient client = new OkHttpClient();
        Request request = new Request.Builder()
                .url("https://api.weixin.qq.com/sns/jscode2session?appid="
                        + appId + "&secret="
                        + appSecret + "&js_code="
                        + code + "&grant_type=authorization_code")
                .build();
        Response response = client.newCall(request).execute();
        if (response.isSuccessful()) {
            LoginDto dto = JSON.parseObject(Objects.requireNonNull(response.body()).byteStream(), LoginDto.class);
            System.out.println(dto.toString());
            //查询是否有当前用户
            Optional<Account> optional = accountRepository.findById(dto.getOpenId());
            if (optional.isPresent()) {
                return optional.get();
            } else {
                Account account = new Account();
                account.setId(dto.getOpenId());
                return accountRepository.save(account);
            }
        }
        return null;
    }

    @Override
    public Account update(Account account) {
        return accountRepository.save(account);
    }

    @Override
    public Account findById(String id) {
        Optional<Account> optional = accountRepository.findById(id);
        return optional.orElse(null);
    }

}
