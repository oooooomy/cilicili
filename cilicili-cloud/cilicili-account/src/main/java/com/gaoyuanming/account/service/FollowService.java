package com.gaoyuanming.account.service;

import com.gaoyuanming.account.model.entity.Follow;

import java.util.List;

public interface FollowService {

    void save(Follow follow) throws Exception;

    List<Follow> findUserFollow(String id);

}
