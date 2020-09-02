package zconf_redis_v7

import (
	"github.com/go-redis/redis/v7"
	"github.com/zionkit/zion/conf"
)

func New(key string) (client *redis.Client, err error) {
	var opts Options
	if err = conf.Load(key, &opts); err != nil {
		return
	}
	client = redis.NewClient(opts.Unwrap())
	return
}

func NewCluster(key string) (client *redis.ClusterClient, err error) {
	var opts ClusterOptions
	if err = conf.Load(key, &opts); err != nil {
		return
	}
	client = redis.NewClusterClient(opts.Unwrap())
	return
}
