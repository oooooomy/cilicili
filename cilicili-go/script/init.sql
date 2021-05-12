CREATE TABLE `t_admin`
(
    `id`       varchar(50) NOT NULL,
    `email`    varchar(50) NOT NULL,
    `password` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_comment`
(
    `id`            varchar(50)  NOT NULL,
    `video_id`      varchar(50)  NOT NULL,
    `user_id`       varchar(50)  NOT NULL,
    `user_nickname` varchar(255) NOT NULL,
    `user_avatar`   varchar(255) NOT NULL,
    `content`       varchar(500) NOT NULL,
    `create_at`     varchar(50)  NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_follow`
(
    `id`           varchar(50) NOT NULL,
    `from_user_id` varchar(50) NOT NULL,
    `to_user_id`   varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_music`
(
    `id`     varchar(50)  NOT NULL,
    `name`   varchar(255) NOT NULL,
    `author` varchar(255) NOT NULL,
    `url`    varchar(255) NOT NULL,
    `poster` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_room`
(
    `id`         varchar(50)  NOT NULL,
    `token`      varchar(255) NOT NULL,
    `count`      int          NOT NULL,
    `title`      varchar(255) NOT NULL,
    `status`     int          NOT NULL,
    `url`        varchar(255) NOT NULL,
    `poster_url` varchar(255) NOT NULL,
    `email`      varchar(255) NOT NULL,
    `is_living`  tinyint      NOT NULL,
    `avatar`     varchar(255) NOT NULL,
    `username`   varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_user`
(
    `id`           varchar(50) NOT NULL,
    `nickname`     varchar(50)  DEFAULT NULL,
    `address`      varchar(255) DEFAULT NULL,
    `avatar_url`   varchar(255) DEFAULT NULL,
    `has_info`     tinyint      DEFAULT '0',
    `gender`       varchar(1)   DEFAULT NULL,
    `follow_count` int          DEFAULT '0',
    `fans_count`   int          DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_video`
(
    `id`            varchar(50)  NOT NULL,
    `user_id`       varchar(50)  NOT NULL,
    `user_nickname` varchar(50)  NOT NULL,
    `user_avatar`   varchar(255) NOT NULL,
    `music_id`      varchar(50)  DEFAULT NULL,
    `title`         varchar(255) DEFAULT NULL,
    `like_count`    int          NOT NULL,
    `comment_count` int          NOT NULL,
    `share_count`   int          NOT NULL,
    `poster`        varchar(255) NOT NULL,
    `url`           varchar(255) NOT NULL,
    `use_music`     tinyint      NOT NULL,
    `music_name`    varchar(255) DEFAULT NULL,
    `music_author`  varchar(255) DEFAULT NULL,
    `music_poster`  varchar(255) DEFAULT NULL,
    `music_url`     varchar(255) DEFAULT NULL,
    `status`        tinyint      NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_video_like`
(
    `id`       varchar(50) NOT NULL,
    `user_id`  varchar(50) NOT NULL,
    `video_id` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;