package person

import "contracts-manager/internal/infrastructure/db/models"

type InsertDTO struct {
	Type  models.PersonType `json:"type" binding:"required,oneof=individual entity"`
	Name  string            `json:"name" binding:"required,min=2,max=255"`
	Code  string            `json:"code" binding:"required,min=3,max=50"`
	Email *string           `json:"email,omitempty" binding:"omitempty,email"`
	Phone *string           `json:"phone,omitempty" binding:"omitempty,e164"`
}

type PersonFilter struct {
	Name  *string
	Type  *string
	Code  *string
	Page  int
	Limit int
}

type PersonListResult struct {
	Data  []models.Person `json:"data"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
	Total int64           `json:"total"`
}

type ImportResponse struct {
	Imported int      `json:"imported"`
	Errors   []string `json:"errors"`
}
