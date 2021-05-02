package com.gaoyuanming.upload.service.impl;

import com.aliyun.oss.OSS;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.PutObjectRequest;
import com.gaoyuanming.common.model.support.ResponseResult;
import com.gaoyuanming.upload.feign.AccountFeign;
import com.gaoyuanming.upload.model.entity.Follow;
import com.gaoyuanming.upload.model.entity.Video;
import com.gaoyuanming.upload.model.entity.VideoLike;
import com.gaoyuanming.upload.repositiry.VideoLikeRepository;
import com.gaoyuanming.upload.repositiry.VideoRepository;
import com.gaoyuanming.upload.service.VideoService;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
public class VideoServiceImpl implements VideoService {

    @Resource
    private AccountFeign accountFeign;

    @Resource
    private VideoRepository videoRepository;

    @Resource
    private VideoLikeRepository videoLikeRepository;

    @Value("${temp-path}")
    private String temp;

    @Value("${ffmpeg-path}")
    private String ffmpeg;

    @Value("${oss.endpoint}")
    private String endpoint;

    @Value("${oss.accessKeyId}")
    private String accessKeyId;

    @Value("${oss.accessKeySecret}")
    private String accessKeySecret;

    @Value("${oss.bucketName}")
    private String bucketName;

    private String putFile(File file) {
        String name = UUID.randomUUID() + ".mp4";
        OSS ossClient = new OSSClientBuilder().build(endpoint, accessKeyId, accessKeySecret);
        PutObjectRequest putObjectRequest = new PutObjectRequest(bucketName, name, file);
        ossClient.putObject(putObjectRequest);
        ossClient.shutdown();
        return "https://" + bucketName + "." + endpoint + "/" + name;
    }

    @Override
    public Video save(Video video) throws Exception {

        String mp4Path = temp + "/" + video.getUrl();

        //mp3合成mp4
        if (video.isUseMusic()) {
            String id = UUID.randomUUID().toString();
            String mp3Path = temp + "/" + video.getMusicUrl();
            String newPath = temp + "/" + id + ".mp4";
            String command = ffmpeg + " -i " + mp4Path + " -i " + mp3Path
                    + " -c:v copy -shortest -c:a aac -strict experimental -map 0:v:0 -map 1:a:0 " + newPath;
            Process exec = Runtime.getRuntime().exec(command);
            exec.waitFor();
            video.setUrl(putFile(new File(newPath)));
        } else {
            video.setUrl(putFile(new File(mp4Path)));
        }

        return videoRepository.save(video);
    }

    @Override
    public void setStatus(String id, boolean status) {
        videoRepository.setStatus(status, id);
    }

    @Override
    public List<Video> findAll() {
        return videoRepository.findAll();
    }

    @Override
    public List<Video> findByUserLike(String id) {
        List<Video> videos = new ArrayList<>();
        List<VideoLike> likes = videoLikeRepository.findAllByUserId(id);
        for (VideoLike vk : likes) {
            Optional<Video> optional = videoRepository.findById(vk.getVideoId());
            optional.ifPresent(videos::add);
        }
        return videos;
    }

    @Override
    public List<Video> findAllByStatus(boolean status) {
        return videoRepository.findAllByStatus(status);
    }

    @Override
    public List<Video> findAllByMusicId(String id) {
        return videoRepository.findAllByMusicId(id);
    }

    @Override
    public List<Video> findAllByFollow(String userId) {
        ResponseResult<List<Follow>> responseResult = accountFeign.findUserFollow(userId);
        if (!responseResult.isStatus()) return null;
        List<Video> videos = new ArrayList<>();
        List<Follow> follows = responseResult.getData();
        for (Follow f : follows) {
            String to = f.getToUserId();
            videos.addAll(findAllByUserId(to));
        }
        return videos;
    }

    @Override
    public List<Video> findAllByUserId(String id) {
        return videoRepository.findAllByUserId(id);
    }

}
