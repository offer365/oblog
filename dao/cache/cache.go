package cache

import (
	"github.com/gomodule/redigo/redis"
	"github.com/offer365/example/redisx"
)

var Conn redis.Conn

func Init(addr,pwd string)  {
	rdb := redisx.NewCache("redis")
	rdb.Init(
		redisx.WithAddr(addr),
		redisx.WithPwd(pwd),
		)
	Conn = rdb.Conn()
}
