package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"mvc/pkg/config"
	"time"
)

type Client struct {
	Client *redis.Pool
	prefix string
}

func NewRedisClient(cfg *config.Redis) *Client {
	client := &redis.Pool{
		MaxIdle:     cfg.MaxIdle,
		MaxActive:   cfg.MaxActive,
		IdleTimeout: time.Second * cfg.Timeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(cfg.ConType, cfg.Host)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			/*if _, err := c.Do("AUTH", redisConf["auth"].(string)); err != nil {
				_ = c.Close()
				fmt.Println(err.Error())
				return nil, err
			}*/
			return c, nil
		},
	}
	redisClient := Client{
		Client: client,
		prefix: cfg.Env,
	}
	return &redisClient
}

func (c *Client) ZAddToRedis(key string, score int64, member interface{}) error {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	_, err := rc.Do("ZADD", key, score, member)
	return err
}

func (c *Client) ZGetAllFromRedis(key string) (interface{}, error) {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	return rc.Do("ZRANGE", key, 0, -1)
}

func (c *Client) SAddToRedisSet(key string, member interface{}) error {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	_, err := rc.Do("SADD", key, member)
	return err
}

func (c *Client) SIsNumberOfRedisSet(key string, member interface{}) (bool, error) {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	value, err := redis.Bool(rc.Do("SISMEMBER", key, member))
	return value, err
}

func (c *Client) GetFromRedis(key string) (interface{}, error) {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	value, err := rc.Do("GET", key)
	return value, err
}

func (c *Client) PutToRedis(key string, value interface{}, timeout int) error {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	_, err := rc.Do("SET", key, value, "EX", timeout)
	return err
}

func (c *Client) DeleteFromRedis(key string) error {
	key = c.appendPrefix(key)
	rc := c.Client.Get()
	defer rc.Close()
	_, err := rc.Do("DEL", key)
	return err
}

func (c *Client) appendPrefix(key string) string {
	return c.prefix + "." + key
}
