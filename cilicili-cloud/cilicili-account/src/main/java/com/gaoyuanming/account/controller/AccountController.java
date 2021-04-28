package com.gaoyuanming.account.controller;

import com.gaoyuanming.account.model.entity.Account;
import com.gaoyuanming.account.service.AccountService;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.io.IOException;

@RestController
public class AccountController {

    @Resource
    private AccountService accountService;

    @GetMapping("/getInfo/{id}")
    public Account getInfo(@PathVariable String id){
        return accountService.findById(id);
    }

    @PutMapping("/update")
    public Account update(@RequestBody Account account) {
        return accountService.update(account);
    }

    @PostMapping("/login")
    public Account login(String code) throws IOException {
        return accountService.login(code);
    }

}
