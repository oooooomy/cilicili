package com.gaoyuanming.upload.controller;

import com.gaoyuanming.upload.model.entity.Comment;
import com.gaoyuanming.upload.service.CommentService;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.List;

@RestController
@RequestMapping("/comment")
public class CommentController {

    @Resource
    private CommentService commentService;

    @PostMapping("")
    public Comment save(@RequestBody Comment comment) {
        return commentService.save(comment);
    }

    @GetMapping("/video/{id}")
    public List<Comment> findByVideo(@PathVariable String id) {
        return commentService.findByVideo(id);
    }

}
