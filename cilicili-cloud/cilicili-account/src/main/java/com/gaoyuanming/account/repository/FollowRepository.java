package com.gaoyuanming.account.repository;

import com.gaoyuanming.account.model.entity.Follow;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface FollowRepository extends JpaRepository<Follow, String> {

    int countAllByFromUserId(String id);

    boolean existsByFromUserIdAndToUserId(String from, String to);

    List<Follow> findAllByFromUserId(String id);
}
