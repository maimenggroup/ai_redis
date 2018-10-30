package ai_redis

import (
	"testing"
)

func TestAISentinel_Init(t *testing.T) {
	aiRds := &AISentinel{}
	var addrs []string
	addrs = append(addrs, "127.0.0.1:28806")
	addrs = append(addrs, "127.0.0.1:28807")
	addrs = append(addrs, "127.0.0.1:28808")

	err := aiRds.Init("mymaster", addrs, "Mindata123", 0)
	if err != nil {
		t.Error("TestAISentinel_Init failed, err ", err.Error())
	} else {
		t.Log("TestAISentinel_Init ok")
	}
	opt := aiRds.Cli.Options()
	t.Log(opt.PoolSize)
	t.Log(opt.PoolTimeout)
}

func TestAISentinel_Rpush(t *testing.T) {
	aiRds := &AISentinel{}
	var addrs []string
	addrs = append(addrs, "127.0.0.1:28806")
	addrs = append(addrs, "127.0.0.1:28807")
	addrs = append(addrs, "127.0.0.1:28808")
	aiRds.Init("mymaster", addrs, "Mindata123", 0)
	aiRds.Delete("jim_test")
	aiRds.Rpush("jim_test", "1")
	aiRds.Rpush("jim_test", "1")
	aiRds.Rpush("jim_test", "1")
	length, err := aiRds.LLen("jim_test")
	if err != nil {
		t.Error("TestAISentinel_Rpush failed, err ", err.Error())
	} else {
		t.Log("queue length ", length)
	}
	aiRds.Lpop("jim_test")
	length, _ = aiRds.LLen("jim_test")
	t.Log("queue length ", length)
}

func TestAISentinel_SetNX(t *testing.T) {
	aiRds := &AISentinel{}
	var addrs []string
	addrs = append(addrs, "127.0.0.1:28806")
	addrs = append(addrs, "127.0.0.1:28807")
	addrs = append(addrs, "127.0.0.1:28808")
	aiRds.Init("mymaster", addrs, "Mindata123", 0)
	res, err := aiRds.SetNX("jim_test", "1", 0)
	if err != nil {
		t.Error("TestAISentinel_Rpush failed, ", err.Error())
	} else {
		t.Log("TestAISentinel_Rpush res ", res)
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

func TestAISentinel_Lpop(t *testing.T) {
	aiRds := &AISentinel{}
	var addrs []string
	addrs = append(addrs, "127.0.0.1:28806")
	addrs = append(addrs, "127.0.0.1:28807")
	addrs = append(addrs, "127.0.0.1:28808")
	aiRds.Init("mymaster", addrs, "Mindata123", 0)
	aiRds.Delete("jim_test")
	result, err := aiRds.Lpop("jim_test")
	if err != nil {
		if err.Error() != "redis: nil" {
			t.Errorf("TestAISentinel_Lpop failed, err [%s]", err.Error())
		}
	} else {
		t.Log("TestAISentinel_Lpop res ", result)
	}
	aiRds.Rpush("jim_test", "hello")
	result, err = aiRds.Lpop("jim_test")
	if err != nil {
		t.Error("TestAISentinel_Lpop failed, err ", err.Error())
	} else {
		t.Log("TestAISentinel_Lpop res ", result)
	}
}


func TestAISentinel_LLen(t *testing.T) {
	aiRds := &AISentinel{}
	var addrs []string
	addrs = append(addrs, "127.0.0.1:28806")
	addrs = append(addrs, "127.0.0.1:28807")
	addrs = append(addrs, "127.0.0.1:28808")
	aiRds.Init("mymaster", addrs, "Mindata123", 0)
	aiRds.Delete("jim_test")
	length, err := aiRds.LLen("jim_test")
	if err != nil {
		t.Error("TestAISentinel_LLen failed, err ", err.Error())
	} else {
		t.Log("queue length ", length)
	}
}

func TestAISentinel_GetBit(t *testing.T) {
	aiRds := &AISentinel{}
	var addrs []string
	addrs = append(addrs, "127.0.0.1:28806")
	addrs = append(addrs, "127.0.0.1:28807")
	addrs = append(addrs, "127.0.0.1:28808")
	aiRds.Init("mymaster", addrs, "Mindata123", 0)
	aiRds.SetBit("hello", 123, 1)
	val, err := aiRds.GetBit("hello", 123)
	if err != nil {
		t.Error("TestAISentinel_GetBit failed, err ", err.Error())
	} else {
		t.Log("TestAISentinel_GetBit ok ", val)
	}
}