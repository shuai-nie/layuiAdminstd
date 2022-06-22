package model

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", //连接地址
		Password: "",               //连接密码
		DB:       0,                //默认连接库
		PoolSize: 100,              //连接池大小
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}

	//set,get示例
	err1 := rdb.Set("name", "yyy", 10000000000).Err() //10s
	if err1 != nil {
		fmt.Println("设置值错误")
		return
	}
	val, err1 := rdb.Get("name").Result()
	if err1 != nil {
		fmt.Println("获取值错误")
		return
	}
	fmt.Println(val)

	//hash hget/hset操作
	// err1 := rdb.HSet("id_1", "name", "yzl").Err()
	// if err1 != nil {
	// 	fmt.Println("设置哈希值错误")
	// 	return
	// }
	// val, err1 := rdb.HGet("id_1", "name").Result()
	// if err1 != nil {
	// 	fmt.Println("获取哈希值错误")
	// 	return
	// }
	// fmt.Println(val)

	//链表lpush/lrange操作
	// err1 := rdb.LPush("id_2", 1, 2, 3).Err()
	// if err1 != nil {
	// 	fmt.Println("设置链表值错误")
	// 	return
	// }
	// val, err1 := rdb.LRange("id_2", 0, -1).Result()
	// if err1 != nil {
	// 	fmt.Println("获取链表值错误")
	// 	return
	// }
	// fmt.Println(val)

	//集合sadd/smembers操作
	// err1 := rdb.SAdd("id_3", 4, 5, 6, 7).Err()
	// if err1 != nil {
	// 	fmt.Println("设置集合值错误")
	// 	return
	// }
	// val, err1 := rdb.SMembers("id_3").Result()
	// if err1 != nil {
	// 	fmt.Println("获取集合值错误")
	// 	return
	// }
	// fmt.Println(val)

	//有序集合zadd/zrangewithscores操作
	// err1 := rdb.ZAdd("id_4", redis.Z{Member: "yzl", Score: 99.0}, redis.Z{Member: "yyy", Score: 97.0}, redis.Z{Member: "zzz", Score: 90.0}).Err()
	// if err1 != nil {
	// 	fmt.Println("设置有序集合值错误")
	// 	return
	// }
	// val, err1 := rdb.ZRangeWithScores("id_4", 0, -1).Result()
	// if err1 != nil {
	// 	fmt.Println("获取有序集合值错误")
	// 	return
	// }
	// fmt.Println(val)
}