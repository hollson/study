-- 该算法先进行范围分片，计算出分片组，组内再求模，综合了范围分片和求模分片的优点。分片组内使用求模可以保证组内的数据分布比较均匀，
-- 分片组之间采用范围分片可以兼顾范围分片的特点。事先规定好分片的数量，数据扩容时按分片组扩容，则原有分片组的数据不需要迁移。由于
-- 分片组内的数据分布比较均匀，所以分片组内可以避免热点数据问题。


CREATE TABLE sharding_by_rang_mod (id bigint null,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (1000, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (1002, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (30000, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (30004, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (40000, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (40005, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (60005, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (60006, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (80006, database());
INSERT INTO `sharding_by_rang_mod` (id,db_nm) VALUES (60008, database());
select * from sharding_by_rang_mod;