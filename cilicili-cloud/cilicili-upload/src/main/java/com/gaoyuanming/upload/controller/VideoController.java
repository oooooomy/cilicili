package com.gaoyuanming.upload.controller;

import com.gaoyuanming.upload.model.entity.Video;
import com.gaoyuanming.upload.model.entity.VideoLike;
import com.gaoyuanming.upload.service.VideoLikeService;
import com.gaoyuanming.upload.service.VideoService;
import com.gaoyuanming.upload.service.impl.VideoLikeServiceImpl;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/video")
public class VideoController {

    @Resource
    private VideoService videoService;

    @Resource
    private VideoLikeService videoLikeService;

    @GetMapping("")
    public List<Video> findAll() {
        return videoService.findAll();
    }

    @PostMapping("/like")
    public void addLike(@RequestBody VideoLike videoLike) throws Exception {
        videoLikeService.save(videoLike);
    }

    @GetMapping("/findByMusicId")
    public List<Video> findByMusicId(String musicId) {
        return videoService.findAllByMusicId(musicId);
    }

    @GetMapping("/findAllByFollow")
    public List<Video> findAllByFollow(String userId) {
        return videoService.findAllByFollow(userId);
    }

    @GetMapping("/findByUserId")
    public List<Video> findByUserId(String userId) {
        return videoService.findAllByUserId(userId);
    }

    @GetMapping("/findByStatus")
    public List<Video> findByStatus(String status) {
        return videoService.findAllByStatus(Boolean.parseBoolean(status));
    }

    @GetMapping("/findUserLikeAndUpload/{id}")
    public Map<String, List<Video>> findUserLikeAndUpload(@PathVariable String id) {
        Map<String, List<Video>> map = new HashMap<>();
        map.put("uploadList", videoService.findAllByUserId(id));
        map.put("likeList", videoService.findByUserLike(id));
        return map;
    }

    @PostMapping("")
    public Video save(@RequestBody Video video) throws Exception {
        video.setStatus(false);
        return videoService.save(video);
    }

    @PutMapping("/{id}/setStatus")
    public void setStatus(@PathVariable String id, String status) {
        System.out.println(status);
        videoService.setStatus(id, Boolean.parseBoolean(status));
    }

}
