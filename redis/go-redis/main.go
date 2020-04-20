package main

import (
	"github.com/go-redis/redis"
	"log"
)



// 声明一个全局的redisClient变量
var redisClient *redis.Client
func main(){
	initClient()
	err:=redisClient.Set("score",100,0).Err()
	if err !=nil{
		log.Print("key set err",err)
	}

	value,err := redisClient.Get("score").Result()
	if err == redis.Nil {
		log.Print("key not exsit")
	} else if err !=nil {
		log.Print(err)
	}
	log.Println(value)
}
// 初始化连接
func initClient() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis-server.imccp:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 集群
func initClient2()(err error){
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}