package cache

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestCache(t *testing.T)  {
	Init("117.107.146.194:7963","qwe!@#123")
	_, err := Conn.Do("set", "a", "11")
	if err != nil {
		t.Fatal("set error,", err)
	}
	str, err := redis.String(Conn.Do("get", "a"))
	t.Log(str, err)

	// 设置哈希表
	_, err = Conn.Do("HMSet", redis.Args{"lisi"}.AddFlat(User{"lisi",11,11})...)
	if err != nil {
		t.Log("hset error,", err)
	}
	r, err := redis.Values(Conn.Do("HGETALL", "lisi"))
	user:=User{}
	redis.ScanStruct(r,&user)
	t.Log(user)

	// 批量写入读取对象(Hashtable)
	// HMSET key field value [field value …]
	// HMGET key field [field …]
	// _, err = Conn.Do("HMSet", "hash1", "a", 101, "b", 101, "c", 102, "d", 103)
	// if err != nil {
	// 	t.Log("hset error,", err)
	// }
	//
	// rs, err := redis.Ints(Conn.Do("HMGet", "hash1", "a", "b", "c", "d"))
	// t.Log(rs)
}

type User struct {
	Name string
	Age int64
	Number int64
}