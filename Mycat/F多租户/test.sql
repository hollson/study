-----------------------建库--------------------------
drop database if exists deeplink_1000011;
create database deeplink_1000011 default charset utf8 collate utf8_general_ci;

drop database if exists deeplink_1000012;
create database deeplink_1000012 default charset utf8 collate utf8_general_ci;

drop database if exists deeplink_1000013;
create database deeplink_1000013 default charset utf8 collate utf8_general_ci;



-----------------------建表--------------------------
-- 创建三张表
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for log_sync
-- ----------------------------
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、keep）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '可用数量',
  `content` varchar(100) DEFAULT NULL COMMENT '内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for tb_device
-- ----------------------------
DROP TABLE IF EXISTS `tb_device`;    ---- tb_device_1 。。。多建几个
CREATE TABLE `tb_device` (
  `id` bigint(20) NOT NULL,  -- 不用自增
  `uid` int(11) NOT NULL COMMENT '渠道ID',
  `serial` char(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`,`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备库信息库';

-- ----------------------------
-- Table structure for tb_hit
-- ----------------------------
DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;

-----------------------测试--------------------------
/*!mycat: schema=deeplink_1000012*/INSERT INTO tb_device(`id`, `uid`, `serial`, `type`) VALUES (next value for MYCATSEQ_DEEPLINK1000012, '0017', 'aaa', 0);

explain
/*!mycat: schema=deeplink_1000012*/ select * from tb_device;


提前预设100个表（0-99）