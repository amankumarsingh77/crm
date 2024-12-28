package auth

import (
	"context"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/amankumarsingh77/cmr/pkg/utils"
	"github.com/gofrs/uuid"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	FindByName(ctx context.Context, name string, query *utils.Pagination) (*models.UsersList, error)
	GetUsers(ctx context.Context, pq *utils.Pagination) (*models.UsersList, error)
}
