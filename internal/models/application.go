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
