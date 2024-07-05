package db

import "github.com/Leepard/yoohob-im-server-lib/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
