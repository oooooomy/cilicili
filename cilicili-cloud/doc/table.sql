DROP
DATABASE IF EXISTS cilicili;
CREATE
DATABASE cilicili;
USE
cilicili;

CREATE TABLE `t_account`
(
    `id`           varchar(50) NOT NULL,
    `nickname`     varchar(255)         DEFAULT NULL,
    `address`      varchar(255)         DEFAULT NULL,
    `avatar_url`   varchar(255)         DEFAULT NULL,
    `has_info`     tinyint     NOT NULL DEFAULT '0',
    `gender`       varchar(1)           DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `t_follow`
(
    `id`           varchar(50) NOT NULL,
    `from_user_id` varchar(50) NOT NULL,
    `to_user_id`   int         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;