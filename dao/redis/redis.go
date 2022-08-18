package redis

// 本文参考的是 go-redis  https://www.liwenzhou.com/posts/Go/go_redis/
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"time"
)

var RDB *redis.Client

// 初始化连接
func InitRedis() (err error) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 100,  // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 超时为 5s
	defer cancel()

	_, err = RDB.Ping(ctx).Result()
	return err
}


// set / get 例子
func redisExample()  {
	ctx := context.Background()
	if err := InitRedis(); err != nil {
		return
	}
	err := RDB.Set(ctx, "key", "value", 0).Err()    // string 的操作 Set
	if err != nil {
		panic(err)
	}

	val, err := RDB.Get(ctx, "key").Result()      // string 的操作 Get
	if err != nil {
		panic(err)
	}
	fmt.Println("value : ", val)

	val2 , err := RDB.Get(ctx, "key2").Result()
	if err == redis.Nil {                // key 不存在
		fmt.Printf("key is not exist, err : %v", err)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("value : ", val2)
	}
}

// zset 例子
func redisExample2() {
	ctx := context.Background()
	if err := InitRedis(); err != nil {
		return
	}
	zsetKey := "language_rank"
	languages := []redis.Z{
		redis.Z{Score : 90.0, Member: "Golang"},
		redis.Z{Score : 98.0, Member: "Java"},
		redis.Z{Score : 95.0, Member: "python"},
		redis.Z{Score : 97.0, Member: "JavaScript"},
		redis.Z{Score : 99.0, Member: "C/C++"},
	}

	// ZADD
	num, err := RDB.ZAdd(ctx,zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed , err: %v\n", err)
		return
	}
	fmt.Printf("zadd %d succ.\n", num)

	// 把 Golang 的分数加 10
	newScore, err := RDB.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Println("newScore==>", newScore)

	// 取出全部的
	result, err := RDB.ZRangeByScoreWithScores(ctx, zsetKey, &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()
	fmt.Println(result)   // [{95 python} {97 JavaScript} {98 Java} {99 C/C++} {100 Golang}]

	// 取出最高的2个

}









