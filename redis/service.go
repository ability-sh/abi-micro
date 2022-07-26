package redis

import (
	"time"

	"github.com/ability-sh/abi-lib/dynamic"
	"github.com/ability-sh/abi-micro/micro"
	R "github.com/go-redis/redis/v8"
)

type redisConfig struct {
	Addr         string `json:"addr"`
	UserName     string `json:"username"`
	Password     string `json:"password"`
	DB           int    `json:"db"`
	PoolSize     int    `json:"pool-size"`
	MinIdleConns int    `json:"min-idle-conns"`
	IdleTimeout  int    `json:"idle-timeout"`
}

type redisService struct {
	config interface{}
	name   string
	client *R.Client
}

func newRedisService(name string, config interface{}) RedisService {
	return &redisService{name: name, config: config}
}

/**
* 服务名称
**/
func (s *redisService) Name() string {
	return s.name
}

/**
* 服务配置
**/
func (s *redisService) Config() interface{} {
	return s.config
}

/**
* 初始化服务
**/
func (s *redisService) OnInit(ctx micro.Context) error {

	var err error = nil
	var cfg = &redisConfig{}

	dynamic.SetValue(cfg, s.config)

	s.client = R.NewClient(&R.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB,
		PoolSize:     cfg.PoolSize,
		Username:     cfg.UserName,
		MinIdleConns: cfg.MinIdleConns,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	})

	return err
}

/**
* 校验服务是否可用
**/
func (s *redisService) OnValid(ctx micro.Context) error {
	return nil
}

func (s *redisService) Client() *R.Client {
	return s.client
}

func (s *redisService) Recycle() {
	if s.client != nil {
		s.client.Close()
		s.client = nil
	}
}
