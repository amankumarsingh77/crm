package models

import (
	"github.com/google/uuid"
	"time"
)

type Application struct {
	ID           uuid.UUID  `json:"id" db:"id" redis:"id" validate:"omitempty"`
	UserID       uuid.UUID  `json:"user_id" db:"user_id" redis:"user_id" validate:"required"`
	CurrentStage string     `json:"current_stage" db:"current_stage" redis:"current_stage" validate:"required,oneof=counselling review approved rejected"`
	Status       string     `json:"status" db:"status" redis:"status" validate:"required,oneof=pending in_progress completed"`
	CollegeName  *string    `json:"college_name,omitempty" db:"college_name" redis:"college_name" validate:"omitempty,lte=255"`
	CourseName   *string    `json:"course_name,omitempty" db:"course_name" redis:"course_name" validate:"omitempty,lte=255"`
	IntakeDate   *time.Time `json:"intake_date,omitempty" db:"intake_date" redis:"intake_date" validate:"omitempty"`
	CreatedAt    time.Time  `json:"created_at,omitempty" db:"created_at" redis:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty" db:"updated_at" redis:"updated_at"`
}

type StageNote struct {
	ID              uuid.UUID `json:"id" db:"id" redis:"id" validate:"omitempty"`
	StageProgressID uuid.UUID `json:"stage_progress_id" db:"stage_progress_id" redis:"stage_progress_id" validate:"required"`
	Note            string    `json:"note" db:"note" redis:"note" validate:"required"`
	CreatedBy       uuid.UUID `json:"created_by" db:"created_by" redis:"created_by" validate:"required"`
	CreatedAt       time.Time `json:"created_at" db:"created_at" redis:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at" redis:"updated_at"`
}

type Document struct {
	ID              uuid.UUID  `json:"id" db:"id" redis:"id" validate:"omitempty"`
	StageProgressID uuid.UUID  `json:"stage_progress_id" db:"stage_progress_id" redis:"stage_progress_id" validate:"required"`
	DocumentType    string     `json:"document_type" db:"document_type" redis:"document_type" validate:"required,lte=50"`
	FileName        string     `json:"file_name" db:"file_name" redis:"file_name" validate:"required,lte=255"`
	S3Path          string     `json:"s3_path" db:"s3_path" redis:"s3_path" validate:"required,lte=512"`
	FileSize        int        `json:"file_size" db:"file_size" redis:"file_size" validate:"required"`
	ContentType     string     `json:"content_type" db:"content_type" redis:"content_type" validate:"required,lte=100"`
	UploadedBy      uuid.UUID  `json:"uploaded_by" db:"uploaded_by" redis:"uploaded_by" validate:"required"`
	IsVerified      bool       `json:"is_verified" db:"is_verified" redis:"is_verified" validate:"required"`
	VerifiedBy      *uuid.UUID `json:"verified_by,omitempty" db:"verified_by" redis:"verified_by" validate:"omitempty"`
	VerifiedAt      *time.Time `json:"verified_at,omitempty" db:"verified_at" redis:"verified_at" validate:"omitempty"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at" redis:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at" redis:"updated_at"`
}

type StageProgress struct {
	ID             uuid.UUID  `json:"id" db:"id" redis:"id" validate:"omitempty"`
	ApplicationID  uuid.UUID  `json:"application_id" db:"application_id" redis:"application_id" validate:"required"`
	Stage          string     `json:"stage" db:"stage" redis:"stage" validate:"required,oneof=counselling college_selection application_status visa loan complete"`
	Status         string     `json:"status" db:"status" redis:"status" validate:"required,oneof=pending in_progress completed"`
	StartDate      time.Time  `json:"start_date" db:"start_date" redis:"start_date" validate:"required"`
	CompletionDate *time.Time `json:"completion_date,omitempty" db:"completion_date" redis:"completion_date" validate:"omitempty"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at" redis:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at" redis:"updated_at"`
}

type Notification struct {
	ID            uuid.UUID  `json:"id" db:"id" redis:"id" validate:"omitempty"`
	UserID        uuid.UUID  `json:"user_id" db:"user_id" redis:"user_id" validate:"required"`
	ApplicationID uuid.UUID  `json:"application_id" db:"application_id" redis:"application_id" validate:"required"`
	Title         string     `json:"title" db:"title" redis:"title" validate:"required,lte=255"`
	Message       string     `json:"message" db:"message" redis:"message" validate:"required"`
	IsRead        bool       `json:"is_read" db:"is_read" redis:"is_read" validate:"required"`
	ReadAt        *time.Time `json:"read_at,omitempty" db:"read_at" redis:"read_at" validate:"omitempty"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at" redis:"created_at"`
}

type AuditLog struct {
	ID         uuid.UUID      `json:"id" db:"id" redis:"id" validate:"omitempty"`
	UserID     uuid.UUID      `json:"user_id" db:"user_id" redis:"user_id" validate:"required"`
	Action     string         `json:"action" db:"action" redis:"action" validate:"required,lte=100"`
	EntityType string         `json:"entity_type" db:"entity_type" redis:"entity_type" validate:"required,lte=50"`
	EntityID   uuid.UUID      `json:"entity_id" db:"entity_id" redis:"entity_id" validate:"required"`
	OldValue   map[string]any `json:"old_value,omitempty" db:"old_value" redis:"old_value" validate:"omitempty"`
	NewValue   map[string]any `json:"new_value,omitempty" db:"new_value" redis:"new_value" validate:"omitempty"`
	IPAddress  string         `json:"ip_address" db:"ip_address" redis:"ip_address" validate:"required,lte=45"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at" redis:"created_at"`
}
