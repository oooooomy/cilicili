package com.gaoyuanming.upload.repositiry;

import com.gaoyuanming.upload.model.entity.VideoLike;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface VideoLikeRepository extends JpaRepository<VideoLike, String> {

    List<VideoLike> findAllByUserId(String id);

    boolean existsByVideoIdAndUserId(String videoId, String userId);

}
