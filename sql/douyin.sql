-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: localhost    Database: db10
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tb_comment`
--

DROP TABLE IF EXISTS `tb_comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_comment` (
  `comment_id` BIGINT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(40) DEFAULT '',
  `content` VARCHAR(40) DEFAULT '',
  `isdeleted` TINYINT DEFAULT '0',
  PRIMARY KEY (`comment_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_favorite`
--

DROP TABLE IF EXISTS `tb_favorite`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_favorite` (
  `favorite_id` BIGINT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(40) DEFAULT '',
  `video_id` BIGINT DEFAULT '0',
  `isdeleted` TINYINT DEFAULT '0',
  PRIMARY KEY (`favorite_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_relation`
--

DROP TABLE IF EXISTS `tb_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_relation` (
  `relation_id` BIGINT NOT NULL AUTO_INCREMENT,
  `follower_id` BIGINT DEFAULT '0',
  `following_id` BIGINT DEFAULT '0',
  `isdeleted` TINYINT DEFAULT '0',
  PRIMARY KEY (`relation_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_user`
--

DROP TABLE IF EXISTS `tb_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_user` (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(40) DEFAULT '',
  `follow_count` INT DEFAULT '0',
  `follower_count` INT DEFAULT '0',
  `is_follow` TINYINT DEFAULT '0',
  `password` CHAR(40) DEFAULT '',
  PRIMARY KEY (`user_id`)
<<<<<<< HEAD
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
=======
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
>>>>>>> 416631de5f0e7587a197f0d1ca9f8bb3dd38d200
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tb_video`
--

DROP TABLE IF EXISTS `tb_video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tb_video` (
<<<<<<< HEAD
  `video_id` BIGINT NOT NULL AUTO_INCREMENT,
  `author_name` VARCHAR(40) DEFAULT '',
  `play_url` VARCHAR(60) DEFAULT '',
  `cover_url` VARCHAR(60) DEFAULT '',
  `favorite_count` INT DEFAULT '0',
  `comment_count` INT DEFAULT '0',
  PRIMARY KEY (`video_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
=======
  `video_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NULL,
  `play_url` varchar(60) DEFAULT '',
  `cover_url` varchar(60) DEFAULT '',
  `favorite_count` int DEFAULT '0',
  `comment_count` int DEFAULT '0',
  `create_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
>>>>>>> 416631de5f0e7587a197f0d1ca9f8bb3dd38d200
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
<<<<<<< HEAD
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
=======
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-16 23:06:26
>>>>>>> 416631de5f0e7587a197f0d1ca9f8bb3dd38d200
