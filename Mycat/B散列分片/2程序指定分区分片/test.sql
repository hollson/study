-- 在程序运行阶段，由程序自主决定路由到哪个分片。

CREATE TABLE sharding_by_substring  (`user_id`  varchar(20) NOT NULL ,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_substring` (user_id,db_nm) VALUES ('05-10000', database());
INSERT INTO `sharding_by_substring` (user_id,db_nm) VALUES ('02-10001', database());
INSERT INTO `sharding_by_substring` (user_id,db_nm) VALUES ('03-10002', database());
select * from `sharding_by_substring`;