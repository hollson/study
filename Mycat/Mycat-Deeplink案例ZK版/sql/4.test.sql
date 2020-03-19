
-- 问题：
-- 1.多个切分多表，如何建立唯一索引
-- 2. GLOBAL.CURID能不能设置为0，即从1开始插入？


-- ======================插入数据======================
-- 1.使用注解选定操作的逻辑库
-- 2.使用mycat的sequence_conf.properties配置自增主键，须大写
/*!mycat: schema=DEEPLINK_1000020*/INSERT INTO tb_device(`id`,`serial`) VALUES (next value for MYCATSEQ_DEEPLINK_1000020, 'AAA');


-- ======================查询数据======================
-- 使用执行计划查看执行过程
explain
/*!mycat: schema=DEEPLINK_1000020*/ select * from tb_device;

-- ======================批量插入======================
/*!mycat:catlet=io.mycat.route.sequence.BatchInsertSequence */insert into user(name) values('Tom'),('Cat'),('Alan');




SET @myArrayOfValue = '2,5,2,23,6,';
WHILE (LOCATE(',', @myArrayOfValue) > 0)
DO
    SET @value = ELT(1, @myArrayOfValue);
    SET @myArrayOfValue= SUBSTRING(@myArrayOfValue, LOCATE(',',@myArrayOfValue) + 1);
    INSERT INTO `EXEMPLE` VALUES(@value, 'hello');
END WHILE;


