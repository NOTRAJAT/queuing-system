package redis_db

import (
	"context"
	"log"
	"queuing_system/env"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	Ctx context.Context
	Rbd redis.Client
}



func InitRedis() *RedisStore{
	 ptr_redis_context:=  &RedisStore{
		Ctx: context.Background(),
		Rbd: *redis.NewClient(
			&redis.Options{
				Addr: env.Env.Redis_Host_PORT,
				Password: env.Env.Redis_Password,
				DB: env.Env.Redis_DB,
			},
		),
	}
	_,err:=ptr_redis_context.Rbd.Ping(ptr_redis_context.Ctx).Result();
	if err!=nil{
		log.Fatal("Redis not connected ",err.Error())
		return nil
	}
	log.Println("Redis successfully Connected at ",env.Env.Redis_Host_PORT)
	return ptr_redis_context
}

