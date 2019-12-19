-- 通过在配置文件中配置可能的枚举id，指定数据分布到不同的物理节点上，本规则适用于按照省份或区县来拆分数据类业务。

CREATE TABLE sharding_by_intfile  (`age`  int NOT NULL ,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_intfile` (age,db_nm) VALUES (10, database());
INSERT INTO `sharding_by_intfile` (age,db_nm) VALUES (11, database());
INSERT INTO `sharding_by_intfile` (age,db_nm) VALUES (12, database());
select * from `sharding_by_intfile`;