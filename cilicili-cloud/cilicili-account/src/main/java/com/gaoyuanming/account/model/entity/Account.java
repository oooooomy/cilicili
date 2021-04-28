package com.gaoyuanming.account.model.entity;

import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Data
@Entity
@Table(name = "t_account")
@NoArgsConstructor
public class Account {

    //用户的微信openId
    @Id
    private String id;

    private String nickname;

    private String address;

    private String gender;

    private String avatarUrl;

    //是否获取过微信的user_info
    private boolean hasInfo;

    private int fansCount;

    private int followCount;

}
