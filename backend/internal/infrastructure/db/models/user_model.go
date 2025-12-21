package models

import "github.com/google/uuid"

type UserType string

const (
	UserTypeRegular UserType = "regular"
	UserTypeAdmin   UserType = "admin"
)

type User struct {
	BaseModel

	Username     string `gorm:"size:50;uniqueIndex" json:"username"`
	Email        string `gorm:"size:100;uniqueIndex" json:"email"`
	PasswordHash string `gorm:"size:255" json:"-"`

	Type UserType `gorm:"type:text;not null;default:'regular';index" json:"type"`

	PersonID *uuid.UUID `gorm:"type:char(36);index" json:"personId"`
	Person   *Person    `gorm:"constraint:OnDelete:SET NULL;foreignKey:PersonID;references:ID"`
}
