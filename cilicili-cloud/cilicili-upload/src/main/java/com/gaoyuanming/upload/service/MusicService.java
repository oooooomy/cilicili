package com.gaoyuanming.upload.service;

import com.gaoyuanming.upload.model.entity.Music;

import java.util.List;

public interface MusicService {

    List<Music> findAll();

    Music save(Music music);

    void delete(String id);

}
