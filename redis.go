package ai_redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type AIRedis struct {
	Host string
	Port int
	Password string
	Db int
	Cli *redis.Client
	ExpireTime int64
}

func (rds *AIRedis)Init(host string, port int, password string, db int, expireTime int64) error {
	rds.Host = host
	rds.Port = port
	rds.Password = password
	rds.Db = db
	rds.ExpireTime = expireTime
	address := fmt.Sprintf("%s:%d", host, port)
	rds.Cli = redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: db,
	})
	_, err := rds.Cli.Ping().Result()
	return err
}

func (rds *AIRedis)Set(key string, value string, expireTime int64) (str string, err error) {
	str, err = rds.Cli.Set(key, value, time.Duration(expireTime)*time.Second).Result()
	if err != nil {
		return
	}
	return
}

func (rds *AIRedis)Delete(key string) error {
	return rds.Cli.Del(key).Err()
}

func (rds *AIRedis)Rpush(key string, value interface{}) (n int64, err error) {
	var v interface{}
	v, err = rds.Cli.RPush(key, value).Result()
	if err != nil {
		return
	}
	n = v.(int64)
	return
}

func (rds *AIRedis)Lpop(key string) (str string, err error) {
	var v interface{}
	v, err = rds.Cli.LPop(key).Result()
	if err != nil {
		return
	}
	str = v.(string)
	return
}

func (rds *AIRedis)LLen(key string) (n int64, err error) {
	var v interface{}
	v, err = rds.Cli.LLen(key).Result()
	if err != nil {
		return
	}
	n = v.(int64)
	return
}

func (rds *AIRedis)SetNX(key string, value string, expireTime int64) (n int64, err error) {
	var ok bool
	ok, err = rds.Cli.SetNX(key, value, time.Duration(expireTime)*time.Second).Result()
	if err != nil {
		return
	}
	if ok {
		n = 1
	} else {
		n = -1
	}
	return
}

func (rds *AIRedis)SetBit(key string, offset int64, value int) (n int64, err error) {
	return rds.Cli.SetBit(key, offset, value).Result()
}

func (rds *AIRedis)GetBit(key string, offset int64) (n int64, err error) {
	return rds.Cli.GetBit(key, offset).Result()
}
