package rpc

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

/*
	使用工厂模式，创造一个UserDao实例
*/
func NewUserDao(address string, maxIdle int, maxActive int, idleTimeout time.Duration) (userDao *UserDao) {
	return &UserDao{
		pool: &redis.Pool{
			MaxIdle:     maxIdle,     //最大空闲的连接数
			MaxActive:   maxActive,   //表示和redis的最大连接数，0表示无限制
			IdleTimeout: idleTimeout, //最大空闲时间
			Dial: func() (redis.Conn, error) { //初始化连接的代码，连接哪个redis
				return redis.Dial("tcp", address)
			},
		},
	}
}

//根据传入的key，返回user实例对象或错误
func (up *UserDao) GetData(key string) (data string, err error) {
	//从redis连接池中获取一个连接
	redisClient := up.pool.Get()
	defer redisClient.Close()
	//通过传入的id去redis查询
	//res, err := redis.String(redisClient.Do("HGet", "users", userId))
	data, err = redis.String(redisClient.Do("Get", key))

	if err != nil {
		return
	}
	return
}
