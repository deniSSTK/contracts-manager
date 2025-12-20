package contract

import (
	"contracts-manager/internal/infrastructure/db/models"
	"time"

	"github.com/google/uuid"
)

type CreateDTO struct {
	Code        string     `json:"code" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	StartDate   *time.Time `json:"startDate"`
	EndDate     *time.Time `json:"endDate"`
}

type UpdateDTO struct {
	Code        *string     `json:"code"`
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

type Filter struct {
	Code        *string
	Title       *string
	Description *string
	Page        int
	Limit       int
}
