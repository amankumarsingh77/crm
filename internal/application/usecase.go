package application

import (
	"context"
	"github.com/amankumarsingh77/cmr/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, application *models.Application) (*models.Application, error)
	Update(ctx context.Context, application *models.Application) (*models.Application, error)
	GetApplicationStatus(ctx context.Context, applicationID string) (string, error)
	GetAllApplications(ctx context.Context) ([]*models.Application, error)
	GetByID(ctx context.Context, applicationID string) (*models.Application, error)
}
