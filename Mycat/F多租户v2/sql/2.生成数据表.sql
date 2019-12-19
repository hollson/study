use deeplink;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device`;
CREATE TABLE `tb_device` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================

use deeplink_1000011;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================

use deeplink_1000012;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================

use deeplink_1000013;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';
DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000014;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000015;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000016;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000017;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000018;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000019;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================
use deeplink_1000020;
DROP TABLE IF EXISTS `log_sync`;
CREATE TABLE `log_sync` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `day` date NOT NULL COMMENT '同步日期',
  `type` varchar(20) DEFAULT NULL COMMENT '上游（uc、cmb等）',
  `count` int(11) DEFAULT '0' COMMENT '总同步数量',
  `hits` int(11) DEFAULT '0' COMMENT '命中数量',
  `content` varchar(100) DEFAULT NULL COMMENT '日志内容',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_device_1`;
CREATE TABLE `tb_device_1` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_2`;
CREATE TABLE `tb_device_2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_3`;
CREATE TABLE `tb_device_3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_4`;
CREATE TABLE `tb_device_4` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_5`;
CREATE TABLE `tb_device_5` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_6`;
CREATE TABLE `tb_device_6` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_7`;
CREATE TABLE `tb_device_7` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_8`;
CREATE TABLE `tb_device_8` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_9`;
CREATE TABLE `tb_device_9` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_device_10`;
CREATE TABLE `tb_device_10` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL COMMENT '设备序列号',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '设备号类型，0:未知，1:原始编号，2：MD5，3:…',
  `create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_serial` (`serial`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='设备信息';

DROP TABLE IF EXISTS `tb_hit`;
CREATE TABLE `tb_hit` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `serial` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- =========================================================