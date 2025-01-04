package application

import (
	"context"
	"github.com/amankumarsingh77/cmr/internal/models"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateApplication(ctx context.Context, tx *sqlx.Tx, application *models.Application) (*models.Application, error)
	CreateStageProgress(ctx context.Context, tx *sqlx.Tx, progress *models.StageProgress) (*models.StageProgress, error)
	CreateDocument(ctx context.Context, tx *sqlx.Tx, document *models.Document) (*models.Document, error)
	CreateNotification(ctx context.Context, tx *sqlx.Tx, notification *models.Notification) (*models.Notification, error)
	CreateStageNote(ctx context.Context, tx *sqlx.Tx, note *models.StageNote) (*models.StageNote, error)
	Update(ctx context.Context, application *models.Application) (*models.Application, error)
	GetByID(ctx context.Context, applicationID string) (*models.Application, error)
	GetByUserID(ctx context.Context, userID string) ([]*models.Application, error)
	//GetApplicationStatus(ctx context.Context, applicationID string) (map[string]string, error)
	//GetApplicationReport(ctx context.Context, applicationID string) (*models.ApplicationReport, error)  Not sure yet
}
