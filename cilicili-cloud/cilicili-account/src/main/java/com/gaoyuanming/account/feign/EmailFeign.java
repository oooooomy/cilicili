package com.gaoyuanming.account.feign;

import com.gaoyuanming.common.model.support.ResponseResult;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

@FeignClient(name = "email-service")
public interface EmailFeign {

    @GetMapping("/check")
    ResponseResult<Object> checkEmailCode(@RequestParam String email, @RequestParam String code);

}
