-- 此规则是截取字符串中的int数值的hash分片。

CREATE TABLE sharding_by_stringhash (ord_no varchar(20) NULL,`db_nm`  varchar(20) NULL);
INSERT INTO `sharding_by_stringhash` (ord_no,db_nm) VALUES (171022237582, database());  
INSERT INTO `sharding_by_stringhash` (ord_no,db_nm) VALUES (171022553756, database());  
select * from sharding_by_stringhash;  