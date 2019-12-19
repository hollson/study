

mysql> create table local_key_test(id int,name varchar(20));


mycat> insert into local_key_test(id,name) values(next value for MYCATSEQ_GLOBAL,@@hostname);
mycat> select * from local_key_test;

-- 自增键配置信息会自动保存
-- 可以使用 mysql>select next value for MYCATSEQ_xxx(自定义的名字); 来查看下一个自增ID