//用户表
CREATE TABLE `tb_user`
(
    `user_id`        bigint NOT NULL AUTO_INCREMENT,
    `name`           varchar(40) DEFAULT '',
    `follow_count`   int         DEFAULT '0',
    `follower_count` int         DEFAULT '0',
    `is_follow`      tinyint     DEFAULT '0',
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci);

//视频表
CREATE TABLE `tb_video`
(
    `video_id`       bigint NOT NULL AUTO_INCREMENT,
    `author_name`    varchar(40) DEFAULT '',
    `play_url`       varchar(60) DEFAULT '',
    `cover_url`      varchar(60) DEFAULT '',
    `favorite_count` int         DEFAULT '0',
    `comment_count`  int         DEFAULT '0',
    PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci);


