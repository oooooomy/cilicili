package com.gaoyuanming.common.utils;

import java.text.SimpleDateFormat;

public final class TimeUtils {

    public static String getNowTimeString(){
        SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        return df.format(System.currentTimeMillis());
    }

}
