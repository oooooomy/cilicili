package com.gaoyuanming.email.cache;

import com.gaoyuanming.email.model.Code;
import org.springframework.stereotype.Component;

import java.util.concurrent.ConcurrentHashMap;

/**
 * 邮箱验证码
 */
@Component
public class CodeCache {

    private final ConcurrentHashMap<String, Code> map;

    public CodeCache() {
        map = new ConcurrentHashMap<>();
    }

    public void put(String key, Code code) {
        map.put(key, code);
    }

    public Code get(String key) {
        return map.get(key);
    }

    public void remove(String key) {
        map.remove(key);
    }

}
