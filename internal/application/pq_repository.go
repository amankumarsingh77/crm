package application

import (
	"context"
	"github.com/amankumarsingh77/cmr/internal/models"
)

type Repository interface {
	Create(ctx context.Context, application *models.Application) (*models.Application, error)
	Update(ctx context.Context, application *models.Application) (*models.Application, error)
	GetByID(ctx context.Context, applicationID string) (*models.Application, error)
	GetByUserID(ctx context.Context, userID string) ([]*models.Application, error)
	//GetApplicationReport(ctx context.Context, applicationID string) (*models.ApplicationReport, error)  Not sure yet
}
