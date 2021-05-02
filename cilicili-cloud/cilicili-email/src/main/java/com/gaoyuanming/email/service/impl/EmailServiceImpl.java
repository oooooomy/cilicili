package com.gaoyuanming.email.service.impl;

import com.gaoyuanming.email.cache.CodeCache;
import com.gaoyuanming.email.model.Code;
import com.gaoyuanming.email.model.Message;
import com.gaoyuanming.email.service.EmailService;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Random;

@Service
public class EmailServiceImpl implements EmailService {

    private final static int EXP = 1000 * 60 * 10;

    @Resource
    private JavaMailSender sender;

    @Value("${spring.mail.username}")
    private String from;

    @Resource
    private CodeCache codeCache;

    @Override
    public void send(Message message) {
        SimpleMailMessage m = new SimpleMailMessage();
        m.setFrom(from);
        m.setTo(message.getSendTo());
        m.setSubject("cilicili");
        m.setText(message.getContent());
        sender.send(m);
    }

    /**
     * 发送验证码
     * @param to 收件人邮箱
     * @throws Exception 异常
     */
    @Override
    public void sendCode(String to) throws Exception {
        Code code = codeCache.get(to);
        if (code != null) {
            //验证码存在且未过期
            if (code.getExp() > System.currentTimeMillis())
                throw new Exception("请勿频繁发送验证码");
        }
        SimpleMailMessage m = new SimpleMailMessage();
        m.setFrom(from);
        m.setTo(to);
        m.setSubject("cilicili 验证码");
        String v = randomCode();
        m.setText("你的验证码为:  " + v + "  十分钟内有效");
        sender.send(m);
        codeCache.put(to, new Code(v, System.currentTimeMillis() + EXP));
    }

    @Override
    public void checkCode(String email, String value) throws Exception {
        Code code = codeCache.get(email);
        if (code == null) throw new Exception("验证码错误，请仔细检查邮箱");
        if (code.getExp() < System.currentTimeMillis()) throw new Exception("验证码已过期，请重新发送");
        if (!code.getValue().equals(value)) throw new Exception("验证码错误，请仔细检查邮箱");
        //登录成功删除code
        codeCache.remove(email);
    }

    private static String randomCode() {
        StringBuilder builder = new StringBuilder();
        Random random = new Random();
        for (int i = 0; i < 6; i++) {
            builder.append(random.nextInt(10));
        }
        return builder.toString();
    }

}
