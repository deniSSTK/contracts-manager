package contract

import (
	"contracts-manager/internal/infrastructure/db/models"
	"time"

	"github.com/google/uuid"
)

type CreateDTO struct {
	Number      string     `json:"number" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
}

type UpdateDTO struct {
	Number      *string     `json:"number"`
	Title       *string     `json:"title"`
	Description **string    `json:"description"`
	StartDate   **time.Time `json:"startDate"`
	EndDate     **time.Time `json:"endDate"`
}

type AddPersonDTO struct {
	ContractID uuid.UUID
	PersonID   uuid.UUID
	Role       models.ContractRole
}
