package redis

import (
	"context"
	"os"
	"time"

	redisotel "github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	redisClient   *redis.Client
	clusterClient *redis.ClusterClient
	isCluster     bool
}

func NewRedisConnection() *RedisConnection {
	redisURL := os.Getenv("REDIS_URL")

	if os.Getenv("REDIS_TYPE") == "cluster" {
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        []string{redisURL},
			PoolSize:     100,
			PoolTimeout:  15 * time.Second,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		})

		// Enable tracing instrumentation.
		if err := redisotel.InstrumentTracing(client); err != nil {
			panic(err)
		}

		return &RedisConnection{
			clusterClient: client,
			isCluster:     true,
		}
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(client); err != nil {
		panic(err)
	}

	return &RedisConnection{
		redisClient:   client,
		clusterClient: nil,
		isCluster:     false,
	}
}

func (con *RedisConnection) Ping(ctx context.Context) *redis.StatusCmd {
	if con.isCluster {
		return con.clusterClient.Ping(ctx)
	}

	return con.redisClient.Ping(ctx)
}

func (c *RedisConnection) Incr(ctx context.Context, key string) *redis.IntCmd {
	if c.isCluster {
		return c.clusterClient.Incr(ctx, key)
	}

	return c.redisClient.Incr(ctx, key)
}

func (con *RedisConnection) Close() {
	if con.isCluster {
		con.clusterClient.Close()
	}

	con.redisClient.Close()
}
