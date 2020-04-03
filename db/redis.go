package db

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var Redgo *redis.Pool

func init(){
	Redgo = newPool()
}

func newPool() *redis.Pool{
	c_file,err_load := LoadConfig("")
	if err_load!=nil{
		log.Println("配置文件加载有误",err_load.Error())
	}
	return &redis.Pool{
		MaxIdle: 3,		//最大空闲连接数
		IdleTimeout: 240 * time.Second,		//空闲连接超时关闭连接
		MaxActive: 0, //最大连接数，0表示不限制
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) {
			c,err:=redis.Dial("tcp", c_file.Redis.Host)
			if err!=nil {
				log.Fatal("redis连接失败",err.Error())
			}
			return c, nil
		},
	}
}