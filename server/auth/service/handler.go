package service

import (
	"log"
	"time"

	"github.com/go-redis/redis/v9"
	"gitlab.cs.washington.edu/assafv/fs3/protos/authservice"
	"gitlab.cs.washington.edu/assafv/fs3/server/auth/config"
)

type AuthServiceHandler struct {
	Logger      *log.Logger
	RedisClient *redis.Client
	authservice.UnimplementedAuthServer
}

func NewAuthServiceHandler(logger *log.Logger) *AuthServiceHandler {
	// lazily create redis connection later
	return &AuthServiceHandler{
		Logger:      logger,
		RedisClient: nil,
	}
}

func (handler *AuthServiceHandler) VerifyRedisClient() {
	if handler.RedisClient != nil {
		return
	}
	handler.RedisClient = redis.NewClient(&redis.Options{
		Addr:            config.GetRedisAddress(),
		Password:        config.GetRedisPassword(),
		DB:              0,
		MaxIdleConns:    100,
		ConnMaxIdleTime: 360 * time.Second,
	})
}
