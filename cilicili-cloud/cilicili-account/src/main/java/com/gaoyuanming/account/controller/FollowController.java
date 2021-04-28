package com.gaoyuanming.account.controller;

import com.gaoyuanming.account.model.entity.Follow;
import com.gaoyuanming.account.service.FollowService;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.List;

@RestController
@RequestMapping("/follow")
public class FollowController {

    @Resource
    private FollowService followService;

    @PostMapping("")
    public void save(@RequestBody Follow follow) throws Exception {
        followService.save(follow);
    }

    @GetMapping("/fromUser/{id}")
    public List<Follow> findUserFollow(@PathVariable String id){
        return followService.findUserFollow(id);
    }

}
