package com.gaoyuanming.account.service.impl;

import com.gaoyuanming.account.model.entity.Account;
import com.gaoyuanming.account.model.entity.Follow;
import com.gaoyuanming.account.repository.AccountRepository;
import com.gaoyuanming.account.repository.FollowRepository;
import com.gaoyuanming.account.service.FollowService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;
import java.util.Optional;

@Service
public class FollowServiceImpl implements FollowService {

    @Resource
    private AccountRepository accountRepository;

    @Resource
    private FollowRepository followRepository;

    @Override
    public void save(Follow follow) throws Exception {
        if (followRepository.existsByFromUserIdAndToUserId(
                follow.getFromUserId(),
                follow.getToUserId())) {
            throw new Exception("已经添加了关注");
        }
        Optional<Account> optional = accountRepository.findById(follow.getToUserId());
        if (optional.isPresent()) {
            Account account = optional.get();
            account.setFansCount(account.getFansCount() + 1);
            accountRepository.save(account);
        }
        optional = accountRepository.findById(follow.getFromUserId());
        if (optional.isPresent()) {
            Account account = optional.get();
            account.setFollowCount(account.getFollowCount() + 1);
            accountRepository.save(account);
        }
        followRepository.save(follow);
    }

    @Override
    public List<Follow> findUserFollow(String id) {
        return followRepository.findAllByFromUserId(id);
    }

}
