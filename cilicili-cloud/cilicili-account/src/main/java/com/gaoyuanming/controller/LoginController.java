package com.gaoyuanming.controller;

import com.gaoyuanming.model.entity.Account;
import com.gaoyuanming.service.AccountService;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
public class LoginController {

    @Resource
    private AccountService accountService;

    @PostMapping("/login")
    public Account login(String code) {
        return accountService.login(code);
    }

}
