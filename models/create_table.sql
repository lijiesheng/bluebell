use bluebell;

--  这里为啥不能用 id 自增的代替user_id
--     原因 1 : 一个新用户注册后，别人通过 user_id , 就知道有多少用户
--     原因 2 : 如果用 id 代替 user_id  [比如，用户是1 ~ 10000],
--             对这个1000 个用户就行分库分表，每个500个，那么id 都是从 1 ~ 500 , 会有重复
--  为啥不用 UUID
    -- 1、作为乱序序列，会严重影响到innodb新行的写入性能。
    -- 2、采用无意义字符串，没有排序。
    -- 3、使用字符串形式存储，数据量大时查询效率比较低。
create table `user`(
                       `id` bigint(20) not null primary key auto_increment,
                       `user_id` bigint(20) not null ,
                       `username` varchar(64) collate utf8mb4_general_ci not null ,
                       `password` varchar(64) collate utf8mb4_general_ci not null ,
                       `email` varchar(64) collate utf8mb4_general_ci,
                       `gender` tinyint(4) not null default 0,
                       `create_time` timestamp not null default CURRENT_TIMESTAMP,
                       `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       unique key  `idx_username` (`username`) USING BTREE ,
                       UNIQUE key  `idx_user_id` (`user_id`) using BTREE
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci

