package ai_redis

import (
	"github.com/go-redis/redis"
	"time"
)

type AISentinel struct {
	Master   string
	Addrs    []string
	Password string
	Db       int
	Cli      *redis.Client
}

func (rds *AISentinel) Init(master string, addrs []string, password string, db int) error {
	rds.Master = master
	rds.Addrs = addrs
	rds.Password = password
	rds.Db = db
	rds.Cli = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    master,
		SentinelAddrs: addrs,
		Password:      password,
		DB:            db,
	})
	_, err := rds.Cli.Ping().Result()
	return err
}

func (rds *AISentinel) Set(key string, value string, expireTime int64) (str string, err error) {
	str, err = rds.Cli.Set(key, value, time.Duration(expireTime)*time.Second).Result()
	if err != nil {
		return
	}
	return
}

func (rds *AISentinel) Delete(key string) error {
	return rds.Cli.Del(key).Err()
}

func (rds *AISentinel) Rpush(key string, value string) (n int64, err error) {
	var v interface{}
	v, err = rds.Cli.RPush(key, value).Result()
	if err != nil {
		return
	}
	n = v.(int64)
	return
}

func (rds *AISentinel) Lpop(key string) (str string, err error) {
	var v interface{}
	v, err = rds.Cli.LPop(key).Result()
	if err != nil {
		return
	}
	str = v.(string)
	return
}

func (rds *AISentinel) LLen(key string) (n int64, err error) {
	var v interface{}
	v, err = rds.Cli.LLen(key).Result()
	if err != nil {
		return
	}
	n = v.(int64)
	return
}

func (rds *AISentinel) SetNX(key string, value string, expireTime int64) (n int64, err error) {
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

func (rds *AISentinel) SetBit(key string, offset int64, value int) (n int64, err error) {
	return rds.Cli.SetBit(key, offset, value).Result()
}

func (rds *AISentinel) GetBit(key string, offset int64) (n int64, err error) {
	return rds.Cli.GetBit(key, offset).Result()
}
