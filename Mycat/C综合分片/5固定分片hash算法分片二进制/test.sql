-- 本条规则类似于十进制的求模运算，区别在于是二进制的操作，是取id的二进制低10位，即id二进制&1111111111。此算法的优点在
-- 于如果按照十进制取模运算，则在连续插入1~10时，1~10会被分到1~10个分片，增大了插入事务的控制难度。而此算法根据二进制则
-- 可能会分到连接的分片，降低了插入事务的控制难度。

CREATE TABLE sharding_by_long (id int(10) null,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (1000, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (1002, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (30000, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (30004, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (4000, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (4000, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (60007, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (60006, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (80006, database());  
INSERT INTO `sharding_by_long` (id,db_nm) VALUES (0, database());  
select * from sharding_by_long;





