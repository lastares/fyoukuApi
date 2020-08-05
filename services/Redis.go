package redisClient

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"time"
)


func Connect()  redis.Conn {

	pool, _ := redis.Dial("tcp", beego.AppConfig.String("redisDb"))
	return pool
}

// 通过连接池获取一个连接
func PoolConnect()  redis.Conn {
	pool := &redis.Pool{
		MaxIdle: 1, // 最大的空闲连接数
		MaxActive: 10, // 最大连接数
		IdleTimeout: 180 * time.Second,
		Wait: true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redisDb"))

			if err != nil {
				return nil, err
			}

			if _, err := c.Do("AUTH", "syf93529"); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
	}
	return pool.Get()
}
