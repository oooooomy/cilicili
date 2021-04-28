package com.gaoyuanming.upload.handler;

import com.gaoyuanming.common.model.support.ResponseResult;
import com.gaoyuanming.upload.annotation.DisableBaseResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@ResponseBody
@RestControllerAdvice
public class GlobalExceptionHandler {

    private final Logger logger = LoggerFactory.getLogger(getClass());

    @DisableBaseResponse
    @ExceptionHandler(value = Exception.class)
    public Object handleException(Exception e) {
        logger.warn(e.getMessage());
        return new ResponseResult<>(400, e.getMessage());
    }

}
