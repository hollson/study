-- 在任意实际库上创建，建议主库
use deeplink;

--  创建MYCAT_SEQUENCE表
DROP TABLE IF EXISTS `MYCAT_SEQUENCE`;
CREATE TABLE MYCAT_SEQUENCE (
`name` VARCHAR(50) NOT NULL,
current_value INT NOT NULL,
increment INT NOT NULL DEFAULT 1,
remark varchar(100),  -- remark 并不是必须的，在这里我是为了让每一个表都对应一个全局的自增，在Remark中配置自增项对应的表名。方便后期维护
PRIMARY KEY(name)) ENGINE=InnoDB;

-- – 获取当前sequence的值（返回当前值,增量）
DROP FUNCTION IF EXISTS `mycat_seq_currval`;
DELIMITER ;;
CREATE DEFINER=`work`@`%` FUNCTION `mycat_seq_currval`(seq_name VARCHAR(50)) RETURNS varchar(64) CHARSET latin1
DETERMINISTIC
BEGIN
DECLARE retval VARCHAR(64);
SET retval="-999999999,null";
SELECT concat(CAST(current_value AS CHAR),",",CAST(increment AS CHAR) ) INTO retval FROM MYCAT_SEQUENCE WHERE name = seq_name;
RETURN retval ;
END
;;
DELIMITER ;

-- 设置sequence值
DROP FUNCTION IF EXISTS `mycat_seq_nextval`;
DELIMITER ;;
CREATE DEFINER=`work`@`%` FUNCTION `mycat_seq_nextval`(seq_name VARCHAR(50)) RETURNS varchar(64) CHARSET latin1
DETERMINISTIC
BEGIN
UPDATE MYCAT_SEQUENCE
SET current_value = current_value + increment WHERE name = seq_name;
RETURN mycat_seq_currval(seq_name);
END
;;
DELIMITER ;

-- 获取下一个sequence值
DROP FUNCTION IF EXISTS `mycat_seq_setval`;
DELIMITER ;;
CREATE DEFINER=`work`@`%` FUNCTION `mycat_seq_setval`(seq_name VARCHAR(50), value INTEGER) RETURNS varchar(64) CHARSET latin1
DETERMINISTIC
BEGIN
UPDATE MYCAT_SEQUENCE
SET current_value = value
WHERE name = seq_name;
RETURN mycat_seq_currval(seq_name);
END
;;
DELIMITER ;


-- 添加自增键信息
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('GLOBAL', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('HIT', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('LOG', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000011', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000012', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000013', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000014', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000015', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000016', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000017', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000018', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000019', 0, 1,'GLOBAL');
INSERT INTO MYCAT_SEQUENCE(name,current_value,increment,remark) VALUES ('DEEPLINK_1000020', 0, 1,'GLOBAL');