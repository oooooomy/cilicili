package com.gaoyuanming.account.model.dto;

import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
public class LoginDto {

    private String sessionKey;

    private String openId;

}
