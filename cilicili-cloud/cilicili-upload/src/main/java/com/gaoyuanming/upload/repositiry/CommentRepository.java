package com.gaoyuanming.upload.repositiry;

import com.gaoyuanming.upload.model.entity.Comment;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface CommentRepository extends JpaRepository<Comment, String> {

    List<Comment> findAllByVideoId(String id);
}
