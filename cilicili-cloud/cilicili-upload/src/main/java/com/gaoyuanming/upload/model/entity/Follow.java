package com.gaoyuanming.upload.model.entity;

import lombok.Data;

@Data
public class Follow {

    private String id;

    //谁关注的
    private String fromUserId;

    //关注的谁
    private String toUserId;

}
