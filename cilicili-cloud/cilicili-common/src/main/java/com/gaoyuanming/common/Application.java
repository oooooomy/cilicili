package com.gaoyuanming.common;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.UUID;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;

public class Application {

    public static void main(String[] args) throws Exception {
        String command = "/Users/soanr/Desktop/ffmpeg -i /Users/soanr/Desktop/test.mp4 -i /Users/soanr/Desktop/IU-Blueming.mp3 -c:v copy -c:a aac -strict experimental /Users/soanr/Desktop/" + UUID.randomUUID().toString() + ".mp4";
        Process process = Runtime.getRuntime().exec(command);
        System.out.println(process);
    }

}
