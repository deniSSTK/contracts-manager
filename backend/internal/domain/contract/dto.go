package contract

import "time"

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
