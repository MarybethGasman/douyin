
DROP TABLE IF EXISTS `tb_comment`;
CREATE TABLE `tb_comment` (
  `comment_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT 0,
  `video_id` bigint DEFAULT 0,
  `content` varchar(40) DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_favorite`;
CREATE TABLE `tb_favorite` (
  `favorite_id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(40) DEFAULT '',
  `video_id` bigint DEFAULT '0',
  `is_deleted` tinyint DEFAULT 0,
  PRIMARY KEY (`favorite_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_relation`;
CREATE TABLE `tb_relation` (
  `relation_id` bigint NOT NULL AUTO_INCREMENT,
  `follower_id` bigint DEFAULT '0',
  `following_id` bigint DEFAULT '0',
  `isdeleted` tinyint DEFAULT '0',
  PRIMARY KEY (`relation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(40) DEFAULT '',
  `follow_count` int DEFAULT '0',
  `follower_count` int DEFAULT '0',
  `is_follow` tinyint DEFAULT '0',
  `password` char(40) DEFAULT '',
  PRIMARY KEY (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_video`;
CREATE TABLE `tb_video` (
    `video_id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) DEFAULT NULL,
    `play_url` varchar(60) CHARACTER SET utf8 DEFAULT '',
    `cover_url` varchar(60) CHARACTER SET utf8 DEFAULT '',
    `favorite_count` int(11) DEFAULT '0',
    `comment_count` int(11) DEFAULT '0',
    `title` text CHARACTER SET utf8 COMMENT '视频标题',
    `create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP) COMMENT '创建时间',
    `update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON
UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;