package com.gaoyuanming.upload.controller;

import com.alibaba.fastjson.JSON;
import com.gaoyuanming.common.model.support.ResponseResult;
import com.gaoyuanming.upload.annotation.DisableBaseResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.ServletOutputStream;
import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.Objects;
import java.util.UUID;

@RestController
@RequestMapping("/file")
public class FileController {

    private final Logger logger = LoggerFactory.getLogger(getClass());

    @Value("${temp-path}")
    private String temp;

    @PostMapping("/{type}")
    @DisableBaseResponse
    public void put(MultipartFile file, @PathVariable String type, HttpServletResponse response) throws Exception {
        if (file == null) throw new Exception("Request parameter missing");
        if (file.isEmpty()) throw new Exception("Request parameter is empty");
        String[] split = Objects.requireNonNull(file.getOriginalFilename()).split("\\.");
        //判断最后一个后缀 防止文件名包含 "."
        if (!type.equals(split[split.length - 1])) throw new Exception("Disallowed file type");
        String id = UUID.randomUUID().toString();
        String path = temp + "/" + id + "." + type;
        file.transferTo(new File(path));

        ServletOutputStream outputStream = response.getOutputStream();
        response.setContentType("UTF-8");
        response.setContentType("application/json");
        ResponseResult<String> result = new ResponseResult<>(id + "." + type);
        outputStream.write(JSON.toJSONString(result).getBytes(StandardCharsets.UTF_8));
    }

    @GetMapping("/mp4/{name}")
    @DisableBaseResponse
    public void getMP4(@PathVariable String name, HttpServletResponse response) throws Exception {
        writeFile(response, name, "video/mp4");
    }

    @GetMapping("/mp3/{name}")
    @DisableBaseResponse
    public void getMP3(@PathVariable String name, HttpServletResponse response) throws Exception {
        writeFile(response, name, "video/mpeg");
    }

    @GetMapping("/image/{name}")
    @DisableBaseResponse
    public void getImg(@PathVariable String name, HttpServletResponse response) throws Exception {
        writeFile(response, name, "image/jpeg");
    }

    private void writeFile(HttpServletResponse response, String name, String type) {
        FileInputStream fileInputStream = null;
        OutputStream outputStream = null;
        try {
            fileInputStream = new FileInputStream(temp + "/" + name);
            byte[] data = new byte[fileInputStream.available()];
            int i = fileInputStream.read(data);
            response.setContentType(type);
            outputStream = response.getOutputStream();
            outputStream.write(data);
            outputStream.flush();
        } catch (IOException e) {
            logger.warn(e.getMessage());
        }
    }


}
