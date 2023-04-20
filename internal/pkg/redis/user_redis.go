package redis

import "context"

type UserRedis struct {
	redisClient *RedisConnection
}

type UserRedisInterface interface {
	IncrUserCount(ctx context.Context) error
}

func NewUserRedis() *UserRedis {
	return &UserRedis{
		redisClient: NewRedisConnection(),
	}
}

func (r *UserRedis) IncrUserCount(ctx context.Context) error {
	const key = "user_count"

	if r.redisClient.isCluster {
		_, err := r.redisClient.clusterClient.Incr(ctx, key).Result()
		return err
	}

	_, err := r.redisClient.redisClient.Incr(ctx, key).Result()
	return err
}
