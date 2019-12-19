CREATE TABLE  sharding_by_month  (create_time   timestamp NULL ON UPDATE CURRENT_TIMESTAMP  ,`db_nm`  varchar(20) NULL);  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2017-10-01', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2017-10-30', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2017-11-11', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2017-11-21', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2017-12-01', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2017-12-31', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2018-01-01', database());  
INSERT INTO sharding_by_month (create_time,db_nm) VALUES ('2018-01-31', database());  
select * from sharding_by_month

-- 注意
-- * schema里的table的dataNode节点数必须：大于rule的开始时间按照分片数计算到现在的个数
-- * 按照自然月计算（无论是28、30、31天都是一个月的）
-- * 分片节点个数可以后增加，但是必须符合第一点说明。