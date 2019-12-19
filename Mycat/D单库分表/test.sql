drop database if exists testdb;
create database testdb default charset utf8 collate utf8_general_ci;

-- 创建三张表
drop table if exists book_1;
create table book_1(
id int(11) not null  PRIMARY KEY,
title varchar(20),
tip char(20) default 'tb1'
);

drop table if exists book_2;
create table book_2(
id int(11) not null  PRIMARY KEY,
title varchar(20),
tip char(20) default 'tb2'
);

drop table if exists book_3;
create table book_3(
id int(11) not null  PRIMARY KEY,
title varchar(20),
tip char(20) default 'tb3'
);


-- mycat 插入数据
-- EXPLAIN
insert into book(id,title) VALUES (next value for MYCATSEQ_BOOK,'java')


