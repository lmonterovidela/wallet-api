package infrastructure

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
	"time"
)

// ICacheProvider is an interface to access a cache
type ICacheProvider interface {
	ConnectCache() (interface{}, error)

	Get(key string) (string, error)

	Set(key string, val interface{}, ttl time.Duration) (string, error)
}

type RedisProvider struct {
	client *redis.Client
}

var instanceCache *redis.Client
var onceCache sync.Once

func (provider *RedisProvider) ConnectCache() (interface{}, error) {

	onceCache.Do(func() {

		// Get config with viper
		cacheHost := fmt.Sprintf("%s:%d", viper.GetString("cache.host"), viper.GetInt("cache.port"))

		client := redis.NewClient(&redis.Options{
			Addr:     cacheHost,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		_, err := client.Ping().Result()
		if err != nil {
			logrus.Errorf("couldn't reach cache: %v", err)
			panic(err)
		}
		logrus.Info("Connected to Redis!")
		instanceCache = client

	})
	return instanceCache, nil
}

func (provider *RedisProvider) Get(key string) (string, error) {
	ret, err := provider.client.Get(key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return ret, err
}

func (provider *RedisProvider) Set(key string, val interface{}, ttl time.Duration) (string, error) {
	return provider.client.Set(key, val, ttl).Result()
}

func NewCacheClient() ICacheProvider {
	provider := &RedisProvider{}
	c, err := provider.ConnectCache()
	if err != nil {
		panic(err)
	}
	provider.client = c.(*redis.Client)
	return provider
}
