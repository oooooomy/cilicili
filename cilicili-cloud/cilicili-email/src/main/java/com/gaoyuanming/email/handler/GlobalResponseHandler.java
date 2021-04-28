package com.gaoyuanming.email.handler;

import com.gaoyuanming.common.model.support.ResponseResult;
import org.jetbrains.annotations.NotNull;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

@RestControllerAdvice(basePackages = "com.gaoyuanming.email.controller")
public class GlobalResponseHandler implements ResponseBodyAdvice<Object> {
    @Override
    public boolean supports(@NotNull MethodParameter methodParameter, @NotNull Class aClass) {
        return true;
    }

    @Override
    public ResponseResult<Object> beforeBodyWrite(Object o, @NotNull MethodParameter methodParameter,
                                                  @NotNull MediaType mediaType, @NotNull Class aClass,
                                                  @NotNull ServerHttpRequest serverHttpRequest,
                                                  @NotNull ServerHttpResponse serverHttpResponse) {
        return new ResponseResult<>(o);
    }
}
