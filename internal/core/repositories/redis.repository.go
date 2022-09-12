package repositories

import (
	"context"
	"github.com/Lenstack/farm_management/internal/utils"
	"github.com/go-redis/redis/v9"
)

type IRedisRepository interface {
	IsTokenBlacklisted(token string) (err error)
	AddTokenToBlacklist(token string) (err error)
}

type RedisRepository struct {
	RedisManager *redis.Client
}

var (
	ctx          = context.Background()
	blackListKey = "Jwt_BlackList"
)

func (rr *RedisRepository) IsTokenBlacklisted(token string) (err error) {
	list, err := rr.RedisManager.SMembersMap(ctx, blackListKey).Result()
	if err != nil {
		return err
	}

	if _, ok := list[token]; ok {
		return utils.TokenInBlackList
	}
	return nil
}

func (rr *RedisRepository) AddTokenToBlacklist(token string) (err error) {
	_, err = rr.RedisManager.SAdd(ctx, blackListKey, token).Result()
	return err
}
