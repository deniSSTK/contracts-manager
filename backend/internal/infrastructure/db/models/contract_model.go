package models

import "time"

type Contract struct {
	BaseModel
	Code        string  `gorm:"size:100; uniqueIndex" json:"code"`
	Title       string  `gorm:"size:255" json:"title"`
	Description *string `gorm:"type:text" json:"description,omitempty"`
	StartDate   *time.Time
	EndDate     *time.Time
}
