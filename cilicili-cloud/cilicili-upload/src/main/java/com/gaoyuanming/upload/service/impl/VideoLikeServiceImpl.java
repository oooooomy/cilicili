package com.gaoyuanming.upload.service.impl;

import com.gaoyuanming.upload.model.entity.Video;
import com.gaoyuanming.upload.model.entity.VideoLike;
import com.gaoyuanming.upload.repositiry.VideoLikeRepository;
import com.gaoyuanming.upload.repositiry.VideoRepository;
import com.gaoyuanming.upload.service.VideoLikeService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.Optional;

@Service
public class VideoLikeServiceImpl implements VideoLikeService {

    @Resource
    private VideoRepository videoRepository;

    @Resource
    private VideoLikeRepository likeRepository;

    @Override
    public void save(VideoLike videoLike) throws Exception {
        if (likeRepository.existsByVideoIdAndUserId(videoLike.getVideoId(), videoLike.getUserId())) {
            throw new Exception("已加入我喜欢");
        }
        Optional<Video> optional = videoRepository.findById(videoLike.getVideoId());
        if (optional.isPresent()) {
            Video video = optional.get();
            video.setLikeCount(video.getLikeCount() + 1);
            videoRepository.save(video);
            likeRepository.save(videoLike);
        }
    }

}
