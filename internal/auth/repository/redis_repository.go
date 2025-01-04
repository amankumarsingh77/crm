package repository

import (
	"context"
	"encoding/json"
	"github.com/amankumarsingh77/cmr/internal/auth"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/go-redis/redis/v8"
	"time"
)

type authRedisRepo struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) auth.RedisRepository {
	return &authRedisRepo{
		redisClient: redisClient,
	}
}

func (r *authRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.User, error) {
	userBytes, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	var user models.User
	if err = json.Unmarshal(userBytes, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRedisRepo) SetUserCtx(ctx context.Context, key string, seconds int, user *models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err = r.redisClient.Set(ctx, key, userBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return err
	}
	return nil
}

func (r *authRedisRepo) DeleteUserCtx(ctx context.Context, key string) error {
	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
