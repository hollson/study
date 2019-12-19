CREATE TABLE  sharding_by_day  (create_time   timestamp NULL ON UPDATE CURRENT_TIMESTAMP  ,`db_nm`  varchar(20) NULL);
INSERT INTO sharding_by_day (create_time,db_nm) VALUES ('2017-10-01', database());
INSERT INTO sharding_by_day (create_time,db_nm) VALUES ('2017-10-10', database());
INSERT INTO sharding_by_day (create_time,db_nm) VALUES ('2017-10-11', database());
INSERT INTO sharding_by_day (create_time,db_nm) VALUES ('2017-10-21', database());
INSERT INTO sharding_by_day (create_time,db_nm) VALUES ('2017-10-31', database());
select * from sharding_by_day;

-- 以上测试结果中有个错误，因为配置的dataNode节点数为3，而2017-10-31为2017-10-01后的第四个10天中的第一天，因此需要至少4个dataNode，节点不够就报如上错误了。