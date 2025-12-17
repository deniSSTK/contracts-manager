package models

import "time"

type Contract struct {
	BaseModel
	Number      string  `gorm:"size:100; uniqueIndex" json:"number"`
	Title       string  `gorm:"size:255" json:"title"`
	Description *string `gorm:"type:text" json:"description,omitempty"`
	StartDate   *time.Time
	EndDate     *time.Time
}
