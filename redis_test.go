package ai_redis

import (
	"testing"
)

func TestAIRedis_Init(t *testing.T) {
	aiRds := &AIRedis{}

	err := aiRds.Init("127.0.1", 21601, "Mindata123", 0)
	if err != nil {
		t.Error("TestAIRedis_Init failed, err ", err.Error())
	} else {
		t.Log("TestAIRedis_Init ok")
	}
	opt := aiRds.Cli.Options()
	t.Log(opt.PoolSize)
	t.Log(opt.PoolTimeout)
}

func TestAIRedis_Rpush(t *testing.T) {
	aiRds := &AIRedis{}
	aiRds.Init("127.0.1", 21601, "Mindata123", 0)
	aiRds.Rpush("jim_test", "1")
	aiRds.Rpush("jim_test", "1")
	aiRds.Rpush("jim_test", "1")
	length, err := aiRds.LLen("jim_test")
	if err != nil {
		t.Error("TestAIRedis_Rpush failed, err ", err.Error())
	} else {
		t.Log("queue length ", length)
	}
	aiRds.Lpop("jim_test")
	length, _ = aiRds.LLen("jim_test")
	t.Log("queue length ", length)
}

func TestAIRedis_SetNX(t *testing.T) {
	aiRds := &AIRedis{}
	aiRds.Init("127.0.1", 21601, "Mindata123", 0)
	res, err := aiRds.SetNX("jim_test", "1", 0)
	if err != nil {
		t.Error("TestAIRedis_Rpush failed, ", err.Error())
	} else {
		t.Log("TestAIRedis_Rpush res ", res)
	}
	aiRds.Delete("jim_test")
	res, err = aiRds.SetNX("jim_test", "1", 0)
	if err != nil {
		t.Error("TestRedisSetX failed, ", err.Error())
	} else {
		t.Log("TestRedisSetX res ", res)
	}
	res, err = aiRds.SetNX("jim_test", "1", 0)
	if err != nil {
		t.Error("TestRedisSetX failed, ", err.Error())
	} else {
		t.Log("TestRedisSetX res ", res)
	}
}

func TestAIRedis_Lpop(t *testing.T) {
	aiRds := &AIRedis{}
	aiRds.Init("127.0.1", 21601, "Mindata123", 0)
	aiRds.Delete("jim_test")
	result, err := aiRds.Lpop("jim_test")
	if err != nil {
		if err.Error() != "redis: nil" {
			t.Errorf("TestAIRedis_Lpop failed, err [%s]", err.Error())
		}
	} else {
		t.Log("TestAIRedis_Lpop res ", result)
	}
	aiRds.Rpush("jim_test", "hello")
	result, err = aiRds.Lpop("jim_test")
	if err != nil {
		t.Error("TestAIRedis_Lpop failed, err ", err.Error())
	} else {
		t.Log("TestAIRedis_Lpop res ", result)
	}
}


func TestAIRedis_LLen(t *testing.T) {
	aiRds := &AIRedis{}
	aiRds.Init("127.0.1", 21601, "Mindata123", 0)
	aiRds.Delete("jim_test")
	length, err := aiRds.LLen("jim_test")
	if err != nil {
		t.Error("TestAIRedis_LLen failed, err ", err.Error())
	} else {
		t.Log("queue length ", length)
	}
}
