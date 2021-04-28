package com.gaoyuanming.upload.service;

import com.gaoyuanming.upload.model.entity.Comment;

import java.util.List;

public interface CommentService {

    Comment save(Comment comment);

    List<Comment> findByVideo(String id);

}
