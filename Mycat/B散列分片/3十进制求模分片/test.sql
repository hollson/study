
CREATE TABLE `sharding_by_mod` (id int(10) null,`db_nm` varchar(20) NULL);
INSERT INTO `sharding_by_mod` (id,db_nm) VALUES (1, database());
INSERT INTO `sharding_by_mod` (id,db_nm) VALUES (2, database());
INSERT INTO `sharding_by_mod` (id,db_nm) VALUES (3, database());
INSERT INTO `sharding_by_mod` (id,db_nm) VALUES (10, database());
select * from sharding_by_mod;