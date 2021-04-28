package com.gaoyuanming.email.model;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class Code {

    private String value;

    private long exp;

}
