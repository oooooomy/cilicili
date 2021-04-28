package com.gaoyuanming.upload.repositiry;

import com.gaoyuanming.upload.model.entity.Video;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.stereotype.Repository;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Repository
public interface VideoRepository extends JpaRepository<Video, String> {

    List<Video> findAllByMusicId(String id);

    List<Video> findAllByUserId(String id);

    List<Video> findAllByStatus(boolean status);

    @Transactional
    @Modifying
    @Query("update Video v set v.status = ?1 where v.id=?2")
    void setStatus(boolean status, String id);

}
