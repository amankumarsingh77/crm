package usecase

import (
	"context"
	"github.com/amankumarsingh77/cmr/config"
	"github.com/amankumarsingh77/cmr/internal/application"
	"github.com/amankumarsingh77/cmr/internal/auth"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/amankumarsingh77/cmr/pkg/logger"
	"github.com/amankumarsingh77/cmr/pkg/utils"
	"github.com/google/uuid"
)

type applicationUC struct {
	cfg     *config.Config
	appRepo application.Repository
	awsRepo auth.AWSRepository
	logger  logger.Logger
}

func NewApplicationUC(cfg *config.Config, appRepo application.Repository, awsRepo auth.AWSRepository, logger logger.Logger) application.UseCase {
	return &applicationUC{
		cfg:     cfg,
		appRepo: appRepo,
		awsRepo: awsRepo,
		logger:  logger,
	}
}

func (u *applicationUC) Create(ctx context.Context, application *models.Application) (*models.Application, error) {
	app, err := u.appRepo.CreateApplication(ctx, application)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (u *applicationUC) Update(ctx context.Context, application *models.Application) (*models.Application, error) {
	if err := utils.ValidateIsOwner(ctx, application.UserID.String()); err != nil {
		return nil, err
	}
	app, err := u.appRepo.Update(ctx, application)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (u *applicationUC) GetApplicationStatus(ctx context.Context, applicationID string) (string, error) {

}

func (u *applicationUC) GetAllApplications(ctx context.Context, adminID uuid.UUID) ([]*models.Application, error) {

}

func (u *applicationUC) GetByID(ctx context.Context, applicationID string) (*models.Application, error) {

}
