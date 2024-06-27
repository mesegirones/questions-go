package cache

import (
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Config interface {
	GetRedisURI() string
}

type CacheRepository struct {
	Client *redis.Client
}

func NewCacheRepository(config Config) *CacheRepository {
	strSplit := strings.Split(config.GetRedisURI(), "/")
	db := 0
	addr := strSplit[0]
	if len(strSplit) > 1 {
		db, _ = strconv.Atoi(strSplit[1])
	}
	redisClient := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     addr,
		DB:       db,
		PoolSize: 100,
	})
	return &CacheRepository{
		Client: redisClient,
	}
}
