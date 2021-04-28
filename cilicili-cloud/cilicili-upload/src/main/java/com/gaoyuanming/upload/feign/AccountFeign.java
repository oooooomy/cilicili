package com.gaoyuanming.upload.feign;

import com.gaoyuanming.common.model.support.ResponseResult;
import com.gaoyuanming.upload.model.entity.Follow;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;

import java.util.List;

@FeignClient(name = "account-service")
public interface AccountFeign {

    @GetMapping("/follow/fromUser/{id}")
    ResponseResult<List<Follow>> findUserFollow(@PathVariable String id);

}
