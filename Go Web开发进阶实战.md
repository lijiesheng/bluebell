## 第 4 章 bluebell 项目实战



### 一：为啥不能用 id 作为用户的user_id

```mysql
--  这里为啥不能用 id 自增的代替user_id
--     原因 1 : 一个新用户注册后，别人通过 user_id , 就知道有多少用户
--     原因 2 : 如果用 id 代替 user_id  [比如，用户是1 ~ 10000],
--             对这个1000 个用户就行分库分表，每个500个，那么id 都是从 1 ~ 500 , 会有重复
create table `user`(
                       `id` bigint(20) not null primary key auto_increment,
                       `user_id` bigint(20) not null ,
                       `username` varchar(64) collate utf8mb4_general_ci not null ,
```



### 二： 分布式 ID 生成器

> 场景：电商促销短时间会有大量的订单涌入到系统中，每秒10w+
>
> ​			明星出轨微博短时间会产生大量微博的评论和转发
>
> 这些业务，在插入数据之前，我们需要将这些订单和消息分配一个唯一的 ID 然后在保存到数据库中。即使我们后端的系统对消息进行了分库分表，也可以按照时间的顺序对这些消息进行排序



> 特点：
>
> 1. 全局唯一性：不能出现有重复的 ID 标识
> 2. 递增型：确保生成的 ID 对于用户或者业务逻辑是递增的
> 3. 高可用性：确保在任何时候都生成正确的ID
> 4. 高性能：在高并发的环境下依然表现良好



#### 一 ： Snowflake 雪花算法

![image-20220816225830761](/Users/doghuang/Library/Application Support/typora-user-images/image-20220816225830761.png)

```markdown
# 第一位：1个bit , 值一直为0
```





