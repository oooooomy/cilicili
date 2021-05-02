package com.gaoyuanming.upload.controller;

import com.gaoyuanming.upload.model.entity.Music;
import com.gaoyuanming.upload.service.MusicService;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.List;

@RestController
@RequestMapping("/music")
public class MusicController {

    @Resource
    private MusicService musicService;

    @PostMapping("")
    public Music save(@RequestBody Music music) {
        music.setName(music.getName().replace("&", " "));
        music.setAuthor(music.getAuthor().replace("&", " "));
        return musicService.save(music);
    }

    @DeleteMapping("{id}")
    public void delete(@PathVariable String id) {
        musicService.delete(id);
    }

    @GetMapping("")
    public List<Music> findAll() {
        return musicService.findAll();
    }

}
