package cache

import (
	"context"
	"encoding/json"
	"github.com/JECSand/go-cqrs-service-boilerplate/pkg/logging"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/config"
	"github.com/JECSand/go-cqrs-service-boilerplate/query_service/core/models"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

const (
	redisUserPrefixKey = "reader:user"
)

type redisCache struct {
	log         logging.Logger
	cfg         *config.Config
	redisClient redis.UniversalClient
}

func NewRedisCache(log logging.Logger, cfg *config.Config, redisClient redis.UniversalClient) *redisCache {
	return &redisCache{
		log:         log,
		cfg:         cfg,
		redisClient: redisClient,
	}
}

func (r *redisCache) PutUser(ctx context.Context, key string, user *models.User) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCache.PutUser")
	defer span.Finish()
	b, err := json.Marshal(user)
	if err != nil {
		r.log.WarnMsg("json.Marshal", err)
		return
	}
	if err = r.redisClient.HSetNX(ctx, r.getRedisUserPrefixKey(), key, b).Err(); err != nil {
		r.log.WarnMsg("redisClient.HSetNX", err)
		return
	}
	r.log.Debugf("HSetNX prefix: %s, key: %s", r.getRedisUserPrefixKey(), key)
}

func (r *redisCache) GetUser(ctx context.Context, key string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "redisCache.GetUser")
	defer span.Finish()
	b, err := r.redisClient.HGet(ctx, r.getRedisUserPrefixKey(), key).Bytes()
	if err != nil {
		if err != redis.Nil {
			r.log.WarnMsg("redisClient.HGet", err)
		}
		return nil, errors.Wrap(err, "redisClient.HGet")
	}
	var user models.User
	if err = json.Unmarshal(b, &user); err != nil {
		return nil, err
	}
	r.log.Debugf("HGet prefix: %s, key: %s", r.getRedisUserPrefixKey(), key)
	return &user, nil
}

func (r *redisCache) DeleteUser(ctx context.Context, key string) {
	if err := r.redisClient.HDel(ctx, r.getRedisUserPrefixKey(), key).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("HDel prefix: %s, key: %s", r.getRedisUserPrefixKey(), key)
}

func (r *redisCache) DeleteAllUsers(ctx context.Context) {
	if err := r.redisClient.Del(ctx, r.getRedisUserPrefixKey()).Err(); err != nil {
		r.log.WarnMsg("redisClient.HDel", err)
		return
	}
	r.log.Debugf("Del key: %s", r.getRedisUserPrefixKey())
}

func (r *redisCache) getRedisUserPrefixKey() string {
	if r.cfg.ServiceSettings.RedisUserPrefixKey != "" {
		return r.cfg.ServiceSettings.RedisUserPrefixKey
	}
	return redisUserPrefixKey
}
