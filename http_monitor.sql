#
# SQL Export
# Created by Querious (1002)
# Created: 2017年06月24日 GMT+8 下午5:17:51
# Encoding: Unicode (UTF-8)
#


DROP TABLE IF EXISTS `usergroup`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `taskip`;
DROP TABLE IF EXISTS `task`;
DROP TABLE IF EXISTS `settings`;
DROP TABLE IF EXISTS `sendlist`;
DROP TABLE IF EXISTS `report`;
DROP TABLE IF EXISTS `fault`;
DROP TABLE IF EXISTS `dailyreport`;


CREATE TABLE `dailyreport` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `uptime_percent` float NOT NULL COMMENT '可用率',
  `uptime_minute` int(5) NOT NULL COMMENT '可用时长',
  `avg_resptime` float NOT NULL COMMENT '平均响应时间',
  `max_resptime` float NOT NULL COMMENT '最大响应时间',
  `min_resptime` float NOT NULL COMMENT '最小响应时间',
  `tid` int(5) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=784 DEFAULT CHARSET=utf8;


CREATE TABLE `fault` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `starttime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `lastchecktime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `respcode` int(5) NOT NULL,
  `outofsize` int(1) NOT NULL DEFAULT '0',
  `tid` int(5) NOT NULL,
  `ip` varchar(20) NOT NULL,
  `isremind` int(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=6255 DEFAULT CHARSET=utf8;


CREATE TABLE `report` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `resptime` float NOT NULL,
  `respcode` int(5) NOT NULL,
  `size` int(10) NOT NULL DEFAULT '0',
  `tid` int(5) NOT NULL,
  `ip` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tid` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=3450 DEFAULT CHARSET=utf8;


CREATE TABLE `sendlist` (
  `sendtype` int(1) NOT NULL COMMENT '0 --email,1 --message',
  `content` varchar(50) NOT NULL,
  `tid` int(5) NOT NULL,
  UNIQUE KEY `content` (`content`,`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `settings` (
  `setting_item` varchar(50) NOT NULL,
  `setting_value` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`setting_item`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `task` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `protocol` varchar(10) NOT NULL,
  `url` varchar(150) NOT NULL,
  `username` varchar(50) NOT NULL COMMENT 'FTP用户名，可为空',
  `password` varchar(50) NOT NULL COMMENT 'FTP密码，可为空',
  `method` varchar(5) NOT NULL COMMENT 'GET/POST',
  `params` varchar(100) NOT NULL COMMENT '请求参数',
  `frequency` int(5) NOT NULL,
  `retry` int(3) NOT NULL,
  `goodcode` int(5) NOT NULL DEFAULT '0',
  `sizerange` varchar(10) NOT NULL,
  `status` int(1) NOT NULL DEFAULT '1' COMMENT '0--monitor pause, 1--monitor start',
  `createtime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `uid` int(5) NOT NULL,
  `gid` int(5) NOT NULL,
  `respbody` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8;


CREATE TABLE `taskip` (
  `tid` int(5) NOT NULL,
  `ip` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `user` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `loginname` varchar(50) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `edit_group_task` int(1) NOT NULL,
  `edit_group_user` int(1) NOT NULL,
  `gid` int(5) NOT NULL,
  `lastlogin` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;


CREATE TABLE `usergroup` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `is_user_admin` int(1) NOT NULL,
  `is_group_admin` int(1) NOT NULL,
  `is_conf_admin` int(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;




SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;
SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


