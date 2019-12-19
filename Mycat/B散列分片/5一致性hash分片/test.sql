-- 一致性hash算法有效解决了分布式数据的扩容问题。因为此规则优点在于扩容时迁移数据量比较少，前提是分片节点比较多，虚拟节点分配多些。
-- 虚拟节点分配的少就会造成数据分布不够均匀。但如果实际分片数据比较少，迁移量也会比较多。

CREATE TABLE sharding_by_murmurhash (id int(10) null,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (1, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (2, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (3, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (4, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (5, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (6, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (7, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (8, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (9, database());
INSERT INTO `sharding_by_murmurhash` (id,db_nm) VALUES (10, database());
select * from sharding_by_murmurhash;