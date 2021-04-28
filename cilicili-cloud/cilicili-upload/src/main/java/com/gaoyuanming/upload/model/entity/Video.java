package com.gaoyuanming.upload.model.entity;

import lombok.Data;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

@Data
@Entity
@Table(name = "t_video")
@NoArgsConstructor
public class Video {

    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "org.hibernate.id.UUIDGenerator")
    private String id;

    private String url;

    private String poster;

    private String userId;

    private String userNickname;

    private String userAvatar;

    private boolean useMusic;

    private String musicId;

    private String musicName;

    private String musicAuthor;

    private String musicPoster;

    private String musicUrl;

    private String title;

    private int likeCount;

    private int commentCount;

    private int shareCount;

    private boolean status;

}
