package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/ability-sh/abi-micro/micro"
	R "github.com/go-redis/redis/v8"
)

type RedisService interface {
	micro.Service
	Client() *R.Client
}

func GetClient(ctx micro.Context, name string) (*R.Client, error) {

	s, err := ctx.GetService(name)

	if err != nil {
		return nil, err
	}

	ss, ok := s.(RedisService)

	if !ok {
		return nil, fmt.Errorf("service %s not instanceof RedisService", name)
	}

	ctx.AddCount("redis", 1)

	return ss.Client(), nil
}

func GetRedis(ctx micro.Context, name string) (Redis, error) {

	client, err := GetClient(ctx, name)

	if err != nil {
		return nil, err
	}

	return &redis{ctx: ctx, client: client, c: context.Background()}, nil
}

type redis struct {
	ctx    micro.Context
	client *R.Client
	c      context.Context
}

func isStatError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(R.Error)
	return !ok
}

func (r *redis) Get(key string) (string, error) {
	st := r.ctx.Step("redis.Get")
	rs, err := r.client.Get(r.c, key).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return rs, err
}

func (r *redis) Set(key string, value string, expiration time.Duration) error {
	st := r.ctx.Step("redis.Set")
	_, err := r.client.Set(r.c, key, value, expiration).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return err
}

func (r *redis) TTL(key string) (time.Duration, error) {
	st := r.ctx.Step("redis.TTL")
	rs, err := r.client.TTL(r.c, key).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return rs, err
}

func (r *redis) Expire(key string, expiration time.Duration) error {
	st := r.ctx.Step("redis.Expire")
	_, err := r.client.Expire(r.c, key, expiration).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return err
}

func (r *redis) Del(key string) error {
	st := r.ctx.Step("redis.Del")
	_, err := r.client.Del(r.c, key).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return err
}

func (r *redis) HSet(key string, itemKey string, value string) error {
	st := r.ctx.Step("redis.HSet")
	_, err := r.client.HSet(r.c, key, itemKey, value).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return err
}

func (r *redis) HGet(key string, itemKey string) (string, error) {
	st := r.ctx.Step("redis.HGet")
	rs, err := r.client.HGet(r.c, key, itemKey).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return rs, err
}

func (r *redis) HDel(key string, itemKey string) error {
	st := r.ctx.Step("redis.HDel")
	_, err := r.client.HDel(r.c, key, itemKey).Result()
	if isStatError(err) {
		st("[err:1] %s", err.Error())
	} else {
		st("")
	}
	return err
}
