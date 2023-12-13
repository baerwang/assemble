package config

import (
	redis "github.com/baerwang/redis-tool"
)

func Init() {
	redisConfig := GetRedis()

	rc := redis.Config{
		Host:        redisConfig.Host,
		Password:    redisConfig.Password,
		DbName:      redisConfig.DbName,
		Prefix:      GetConfig().Project.Name,
		Port:        redisConfig.Port,
		MaxIdle:     redisConfig.MaxIdle,
		IdleTimeout: redisConfig.IdleTimeOut,
	}

	if err := redis.LoadRedisSession(rc); err != nil {
		panic("init redis fail : " + err.Error())
	}
}
