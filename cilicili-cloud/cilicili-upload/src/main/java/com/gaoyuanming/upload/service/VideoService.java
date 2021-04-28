package com.gaoyuanming.upload.service;

import com.gaoyuanming.upload.model.entity.Video;

import java.util.List;

public interface VideoService {

    Video save(Video video) throws Exception;

    void setStatus(String id, boolean status);

    List<Video> findAll();

    List<Video> findByUserLike(String id);

    List<Video> findAllByStatus(boolean status);

    List<Video> findAllByMusicId(String id);

    List<Video> findAllByFollow(String userId);

    List<Video> findAllByUserId(String id);

}
