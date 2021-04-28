package com.gaoyuanming.email.controller;

import com.gaoyuanming.common.model.support.ResponseResult;
import com.gaoyuanming.email.model.Message;
import com.gaoyuanming.email.service.EmailService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
public class EmailController {

    @Resource
    private EmailService emailService;

    @PostMapping("/simple")
    public void send(@RequestBody Message message) {
        emailService.send(message);
    }

    @GetMapping("/code")
    public void sentCodeEmail(String sendTo) throws Exception {
        emailService.sendCode(sendTo);
    }

    @GetMapping("/check")
    public void check(String email, String code) throws Exception {
        emailService.checkCode(email, code);
    }

}
