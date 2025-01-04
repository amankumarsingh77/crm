package respository

import (
	"context"
	"fmt"
	"github.com/amankumarsingh77/cmr/internal/application"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/jmoiron/sqlx"
)

type applicationRepo struct {
	db *sqlx.DB
}

func NewApplicationRepository(db *sqlx.DB) application.Repository {
	return &applicationRepo{
		db: db,
	}
}

func (r *applicationRepo) CreateApplication(ctx context.Context, tx *sqlx.Tx, application *models.Application) (*models.Application, error) {
	a := &models.Application{}
	err := tx.QueryRowxContext(
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

func (r *applicationRepo) CreateStageProgress(ctx context.Context, tx *sqlx.Tx, progress *models.StageProgress) (*models.StageProgress, error) {
	s := &models.StageProgress{}
	if err := tx.QueryRowxContext(
		ctx,
		createStageProgress,
		&progress.ApplicationID,
		&progress.Stage,
		&progress.Status,
		&progress.StartDate,
	).StructScan(s); err != nil {
		return nil, err
	}
	return s, nil
}

func (r *applicationRepo) CreateDocument(ctx context.Context, tx *sqlx.Tx, document *models.Document) (*models.Document, error) {
	d := &models.Document{}
	if err := tx.QueryRowxContext(
		ctx,
		createDocument,
		&document.StageProgressID,
		&document.DocumentType,
		&document.FileName,
		&document.S3Path,
		&document.FileSize,
	).StructScan(d); err != nil {
		return nil, err
	}
	return d, nil
}

func (r *applicationRepo) CreateNotification(ctx context.Context, tx *sqlx.Tx, notification *models.Notification) (*models.Notification, error) {
	n := &models.Notification{}
	if err := tx.QueryRowxContext(
		ctx,
		createNotification,
		&notification.UserID,
		&notification.ApplicationID,
		&notification.Title,
		&notification.Message,
		&notification.IsRead,
	).StructScan(n); err != nil {
		return nil, err
	}
	return n, nil
}

func (r *applicationRepo) CreateStageNote(ctx context.Context, tx *sqlx.Tx, note *models.StageNote) (*models.StageNote, error) {
	sn := &models.StageNote{}
	if err := tx.QueryRowxContext(
		ctx,
		createStageNote,
		&note.StageProgressID,
		&note.Note,
		&note.CreatedBy,
	).StructScan(sn); err != nil {
		return nil, err
	}
	return sn, nil
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
