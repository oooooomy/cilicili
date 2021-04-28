package com.gaoyuanming.upload.handler;

import com.gaoyuanming.common.model.support.ResponseResult;
import com.gaoyuanming.upload.annotation.DisableBaseResponse;
import org.jetbrains.annotations.NotNull;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

@ControllerAdvice(basePackages = "com.gaoyuanming.upload.controller")
public class GlobalResponseHandler implements ResponseBodyAdvice<Object> {
    @Override
    public boolean supports(@NotNull MethodParameter methodParameter, @NotNull Class aClass) {
        //如果方法上带有DisableBaseResponse注解， 不处理返回false
        return !methodParameter.hasMethodAnnotation(DisableBaseResponse.class);
    }

    @Override
    public ResponseResult<Object> beforeBodyWrite(Object o, @NotNull MethodParameter methodParameter,
                                                  @NotNull MediaType mediaType, @NotNull Class aClass,
                                                  @NotNull ServerHttpRequest serverHttpRequest,
                                                  @NotNull ServerHttpResponse serverHttpResponse) {
        return new ResponseResult<>(o);
    }
}
