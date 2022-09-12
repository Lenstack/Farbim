package services

import (
	"github.com/Lenstack/farm_management/internal/core/repositories"
	"github.com/go-redis/redis/v9"
)

type IRedisService interface {
	IsTokenBlacklisted(token string) (err error)
	AddTokenToBlacklist(token string) (err error)
}

type RedisService struct {
	redisRepository repositories.RedisRepository
	redisManager    *redis.Client
}

func NewRedisService(redisManager *redis.Client) *RedisService {
	return &RedisService{
		redisRepository: repositories.RedisRepository{
			RedisManager: redisManager,
		},
	}
}

func (rs *RedisService) IsTokenBlacklisted(token string) (err error) {
	return rs.redisRepository.IsTokenBlacklisted(token)
}

func (rs *RedisService) AddTokenToBlacklist(token string) (err error) {
	return rs.redisRepository.AddTokenToBlacklist(token)
}
