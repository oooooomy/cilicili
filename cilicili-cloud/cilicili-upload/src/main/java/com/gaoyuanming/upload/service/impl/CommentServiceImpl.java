package com.gaoyuanming.upload.service.impl;

import com.gaoyuanming.common.utils.TimeUtils;
import com.gaoyuanming.upload.model.entity.Comment;
import com.gaoyuanming.upload.model.entity.Video;
import com.gaoyuanming.upload.repositiry.CommentRepository;
import com.gaoyuanming.upload.repositiry.VideoRepository;
import com.gaoyuanming.upload.service.CommentService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;
import java.util.Optional;

@Service
public class CommentServiceImpl implements CommentService {

    @Resource
    private CommentRepository commentRepository;

    @Resource
    private VideoRepository videoRepository;


    @Override
    public Comment save(Comment comment) {
        Optional<Video> optional = videoRepository.findById(comment.getVideoId());
        if (optional.isPresent()) {
            Video video = optional.get();
            video.setCommentCount(video.getCommentCount() + 1);
            videoRepository.save(video);
        }

        comment.setCreateAt(TimeUtils.getNowTimeString());
        commentRepository.save(comment);
        return comment;
    }

    @Override
    public List<Comment> findByVideo(String id) {
        return commentRepository.findAllByVideoId(id);
    }

}
