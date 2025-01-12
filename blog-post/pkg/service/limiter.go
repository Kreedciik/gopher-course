package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Limiter interface {
	CheckRequest(userId string, ctx *gin.Context) error
}

type LimiterService struct {
	rdb *redis.Client
}

func NewLimiterService(rdb *redis.Client) *LimiterService {
	return &LimiterService{
		rdb,
	}
}

func (l *LimiterService) CheckRequest(userId string, ctx *gin.Context) error {
	const requestLimit = 10
	set, err := l.rdb.SetNX(ctx, userId, 1, time.Minute).Result()

	if err != nil {
		return err
	}

	if !set {
		count, err := l.rdb.Incr(ctx, userId).Result()
		if err != nil {
			return err
		}
		if count > requestLimit {
			ttl, _ := l.rdb.TTL(ctx, userId).Result()
			return fmt.Errorf("your request is limited, please make request after %d seconds", ttl)
		}
	} else {
		return nil
	}

	return nil
}
