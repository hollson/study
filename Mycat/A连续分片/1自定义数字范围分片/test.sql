CREATE TABLE  auto_sharding_long  (`age`  int NOT NULL ,`db_nm`  varchar(20) NULL);
INSERT INTO auto_sharding_long (age,db_nm) VALUES (20000, database());
INSERT INTO auto_sharding_long (age,db_nm) VALUES (25000, database());
INSERT INTO auto_sharding_long (age,db_nm) VALUES (35000, database());
select * from auto_sharding_long;