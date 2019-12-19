-- 单月内按照小时拆分，最小粒度是小时，一天最多可以有24个分片，最少1个分片，下个月从头开始循环，每个月末需要手工清理数据。
-- 字段为字符串类型，yyyymmddHH 10位。

CREATE TABLE  sharding_by_hour  (create_time   timestamp NULL ON UPDATE CURRENT_TIMESTAMP  ,`db_nm`  varchar(20) NULL,sharding_col varchar(10) null);
INSERT INTO sharding_by_hour (sharding_col,create_time,db_nm) VALUES ('2017100101','2017-10-01', database());
INSERT INTO sharding_by_hour (sharding_col,create_time,db_nm) VALUES ('2017100108','2017-10-01', database());
INSERT INTO sharding_by_hour (sharding_col,create_time,db_nm) VALUES ('2017100109','2017-10-01', database());
INSERT INTO sharding_by_hour (sharding_col,create_time,db_nm) VALUES ('2017100116','2017-10-01', database());
INSERT INTO sharding_by_hour (sharding_col,create_time,db_nm) VALUES ('2017100117','2017-10-01', database());
INSERT INTO sharding_by_hour (sharding_col,create_time,db_nm) VALUES ('2017100123','2017-10-01', database());
select * from sharding_by_hour;