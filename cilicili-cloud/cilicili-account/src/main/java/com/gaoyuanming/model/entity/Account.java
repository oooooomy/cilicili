package com.gaoyuanming.model.entity;

import lombok.Data;

@Data
public class Account {

    private String id;

    //微信用户的openId
    private String openId;

    private String nickname;

    private String description;

}
