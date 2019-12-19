
-- 其思想与范围求模一致，由于日期取模方法会出现数据热点问题，所以先根据日期分组，再根据时间hash使得短期内数据分布得更均匀。
-- 其优点是可以避免扩容时的数据迁移，又可以在一定程度上避免范围分片的热点问题，要求日期格式尽量精确，不然达不到局部均匀的目的。

CREATE TABLE `sharding_by_date_hash` (`create_time` datetime NOT NULL,`db_nm` varchar(20) NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-01-12 00:01:02', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-01-21 01:02:09', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-01-28 12:00:12', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-01-02 11:00:00', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-02-26 10:00:09', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-03-01 22:01:02', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-02-02 17:09:08', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-02-23 11:00:04', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-02-24 18:12:09', database());
INSERT INTO `sharding_by_date_hash` (create_time,db_nm) VALUES ('2018-02-21 07:12:00', database());
select * from sharding_by_date_hash;