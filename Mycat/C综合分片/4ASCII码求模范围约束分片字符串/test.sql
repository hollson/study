-- 此种规则类似于取模范围约束，此规则支持数据符号字母取模。

CREATE TABLE sharding_by_ascii (id varchar(20) null,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES ("1000a", database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES ("1002A", database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES (30000, database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES (30004, database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES ("4000B", database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES ("4000b", database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES (60007, database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES (60006, database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES (80006, database());
INSERT INTO `sharding_by_ascii` (id,db_nm) VALUES ("abcd0", database());
select * from sharding_by_ascii;