package person

import "contracts-manager/internal/infrastructure/db/models"

type CreateDTO struct {
	Type  models.PersonType `json:"type" binding:"required,oneof=individual entity"`
	Name  string            `json:"name" binding:"required,min=2,max=255"`
	Code  string            `json:"code" binding:"required,min=3,max=50"`
	Email *string           `json:"email,omitempty" binding:"omitempty,email"`
	Phone *string           `json:"phone,omitempty" binding:"omitempty,e164"`
}

type UpdateDTO struct {
	Type  *models.PersonType `json:"type,omitempty"`
	Name  *string            `json:"name,omitempty"`
	Code  *string            `json:"code,omitempty"`
	Email **string           `json:"email,omitempty"`
	Phone **string           `json:"phone,omitempty"`
}

type Filter struct {
	Name  *string
	Type  *string
	Code  *string
	Page  int
	Limit int
}

type ImportResponse struct {
	Imported int      `json:"imported"`
	Errors   []string `json:"errors"`
}
