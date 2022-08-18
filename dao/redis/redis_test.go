package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-redis/redis/v9"
	"testing"
)

func TestRedis(t *testing.T) {
	err := InitRedis()
	if err != nil {
		fmt.Println("redis 连接错误", err)
	} else {
		fmt.Println("redis 连接成功")
	}
}


func TestV8Example(t *testing.T) {
	redisExample()
}

func TestRedisExample2(t *testing.T)  {
	redisExample2()
}

// 关于 keys 复制 commands_test Describe("keys", func() {})
func TestKeys(t *testing.T) {
	ctx := context.Background()
	err := InitRedis()

	// 删除 Del，使用 del 删除 大key 会造成长时间的阻塞，甚至崩溃
	// 大key 是指 key 的 value 是个庞然大物，如 Hashes, Sorted Sets, Lists, Sets
	// 日积月累之后，会变得非常大，直接使用 del 进行删除，会导致长时间的阻塞，甚至崩溃
	// Redis 是单线程的，单个命令执行时间长了会阻塞其他命令，容易引起雪崩
	// 非字符串的 bigkey , 不要使用 del 删除。 可以使用 unlink
	err = RDB.Set(ctx, "key1", "Hello", 0).Err()
	occur(err, "redis set failed")
	err = RDB.Set(ctx, "key2", "world", 0).Err()
	occur(err, "redis set failed")

	num, err := RDB.Del(ctx, "key1", "key2", "key3").Result()
	occur(err, "redis del failed")
	fmt.Printf("del num is %d\n", num)


	// 删除 Unlink 将实际的键和键空间断开连接，实际的删除将会异步进行。不会阻塞
	// 场景 删除 大 key 的时候使用，并且占用空间积累速度别是特别快
	// 工作原理：
	//  1、把所有的命名空间 key 删除，立即返回，不阻塞
	//  2、后台线程执行真正的释放空间
	err = RDB.Set(ctx, "key1", "Hello", 0).Err()
	occur(err, "redis set failed")
	err = RDB.Set(ctx, "key2", "world", 0).Err()
	occur(err, "redis set failed")

	num, err = RDB.Unlink(ctx, "key1", "key2", "key2").Result()
	occur(err, "redis del failed")
	fmt.Printf("unlink num is %d\n", num)



	// 序列化给定的 key 为字节 Dump , key 需要在 redis 中存在
	err = RDB.Set(ctx, "key", "Hello", 0).Err()
	occur(err, "redis set failed")

	dump := RDB.Dump(ctx, "key")
	fmt.Printf("%v\n", dump)

	// 存在 Exists
	err = RDB.Set(ctx, "key", "Hello", 0).Err()
	n, err := RDB.Exists(ctx, "key1").Result()
	fmt.Printf("key num is %d\n", n)  // 1
	n, err = RDB.Exists(ctx, "key2").Result()
	fmt.Printf("key num is %d\n", n)   // 0

	n, err = RDB.Exists(ctx, "key2", "key1").Result()
	fmt.Printf("key num is %d\n", n)   // 0

	// 过期

	//

}

func occur(err error, str string) {
	if err != nil {
		fmt.Println(str, err)
		panic(err)
	}
}

// 关于 strings


// 关于 hashes


// 关于 lists

// 关于 sets

// 关于 sorted sets


// 关于 scanning
func TestScanning(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	for i := 0; i < 1000; i++ {
		RDB.Set(ctx, fmt.Sprintf("key%d", i), fmt.Sprintf("hello+%d", i), 0)
	}
	// 参数 ctx
	// 参数 cursor  游标
	// 参数 match   匹配的模式
	// 参数 count   指定每次遍历多少个集合
	result, cursor, _ := RDB.Scan(ctx, 0, "key56*", 1000).Result()
	fmt.Println(result)
	fmt.Println(cursor)
}

func TestScanType(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	// 参数 ctx
	// 参数 cursor 游标
	// 参数 match 匹配的模式
	// 参数 count 指定每次遍历多少个集合
	// 参数 keyType
	result, cursor, _ := RDB.ScanType(ctx, 0, "key12*", 1000, "string").Result()
	fmt.Println(result)
	fmt.Println(cursor)
}

// 查找 set 集合的 key
func TestSScan(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	for i := 0; i < 1000; i++ {
		RDB.SAdd(ctx, "myset", fmt.Sprintf("member%d", i))
	}
	result, cursor, _ := RDB.SScan(ctx, "myset", 0, "member12*", 1000).Result()
	fmt.Println(result)
	fmt.Println(cursor)
}

// 查找 hase 集合的 key
func TestHScan(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	for i := 0; i < 1000; i++ {
		 RDB.HSet(ctx, "myhash", fmt.Sprintf("key%d", i), "hello")
	}
	result, cursor, _ := RDB.HScan(ctx, "myhash", 0, "key12*", 1000).Result()
	fmt.Println(result)
	fmt.Println(cursor)
}

// 查找 zset 集合的 key
func TestZScan(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	for i := 0; i < 1000; i++ {
		RDB.ZAdd(ctx, "z_myset", redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("member%d", i),
		})
	}
	result, cursor, _ := RDB.ZScan(ctx, "z_myset", 0, "member12*", 1000).Result()
	fmt.Println(result)
	fmt.Println(cursor)
}

func TestFlushDB(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	RDB.FlushDB(ctx)
}


// ************** string ******************

func TestStingsAppend(t *testing.T) {
	ctx := context.Background()
	InitRedis()
	result, _ := RDB.Append(ctx, "key1", "value1").Result()
	_, _ = RDB.Append(ctx, "key1", "value1").Result()
	fmt.Println(result)
}

func TestStringsBitCount (t *testing.T) {

}






