package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
	"time"
	. "user-srv/utils"
)

var RedisCnn *redis.Client

func InitRedis(){
	RedisCnn = redis.NewClient(&redis.Options{
		Addr:     YamlConfig.Redis.Addr,
		Password: YamlConfig.Redis.Password, // no password set
		DB:       YamlConfig.Redis.Db,  // use default DB
	})

	_ , err := RedisCnn.Ping().Result()
	if err != nil{
		log.Fatalln("conncet redis error: ", err)
	}
}

func Set(Key, Value string, Expira time.Duration) error{
	err := RedisCnn.Set(Key, Value, Expira).Err()
	if err != nil{
		return err
	}else{
		return nil
	}
}

func Get(Key string) (string ,error){
	val, err := RedisCnn.Get(Key).Result()
	if err == redis.Nil {
		err = errors.New("key does not exist")
		return "", err
	} else if err != nil {
		return "", err
	} else {
		return val, nil
	}
}

func Delete(Key string) error{
	err := RedisCnn.Del(Key).Err()
	if err != nil{
		return err
	}else{
		return nil
	}
}

func GetEmailCode() string{
	keys := RedisCnn.Keys("*")
	keysArray, _ := keys.Result()
	code := keysArray[len(keysArray)-1]
	return code
}