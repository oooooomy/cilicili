package com.gaoyuanming.upload.service.impl;

import com.gaoyuanming.upload.model.entity.Music;
import com.gaoyuanming.upload.repositiry.MusicRepository;
import com.gaoyuanming.upload.service.MusicService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Service
public class MusicServiceImpl implements MusicService {

    @Resource
    private MusicRepository musicRepository;

    @Override
    public List<Music> findAll() {
        return musicRepository.findAll();
    }

    @Override
    public Music save(Music music) {
        return musicRepository.save(music);
    }

    @Override
    public void delete(String id) {
        musicRepository.deleteById(id);
    }

}
