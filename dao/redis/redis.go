package redis

// 本文参考的是 go-redis  https://www.liwenzhou.com/posts/Go/go_redis/
import "github.com/go-redis/redis"

var rdb *redis.Client

// 初始化连接
func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
