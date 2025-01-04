package respository

import (
	"context"
	"fmt"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/jmoiron/sqlx"
)

type applicationRepo struct {
	db *sqlx.DB
}

func (r *applicationRepo) Create(ctx context.Context, application *models.Application) (*models.Application, error) {
	a := &models.Application{}
	err := r.db.QueryRowxContext(
		ctx,
		createApplication,
		&application.UserID,
		&application.CollegeName,
		&application.CourseName,
		&application.IntakeDate,
		&application.CurrentStage,
		&application.Status,
	).StructScan(a)
	if err != nil {
		return nil, fmt.Errorf("failed to create application: %v", err)
	}
	return a, nil
}

func (r *applicationRepo) Update(ctx context.Context, application *models.Application) (*models.Application, error) {
	a := &models.Application{}
	if err := r.db.GetContext(
		ctx,
		a,
		updateApplication,
		&application.CollegeName,
		&application.CourseName,
		&application.IntakeDate,
		&application.CurrentStage,
		&application.Status,
		&application.UpdatedAt,
		&application.ID,
	); err != nil {
		return nil, fmt.Errorf("failed to update application : %v", err)
	}
	return a, nil
}

func (r *applicationRepo) GetByID(ctx context.Context, applicationID string) (*models.Application, error) {
	a := &models.Application{}
	if err := r.db.GetContext(ctx, a, getApplicationByID, applicationID); err != nil {
		return nil, fmt.Errorf("failed to get application by id: %v", err)
	}
	return a, nil
}

func (r *applicationRepo) GetByUserID(ctx context.Context, userID string) ([]*models.Application, error) {
	var applications []*models.Application
	if err := r.db.SelectContext(ctx, &applications, getApplicationByUserID, userID); err != nil {
		return nil, fmt.Errorf("failed to get application by user id: %v", err)
	}
	return applications, nil
}
