package models

import "github.com/google/uuid"

type ContractRole string

const (
	RoleSignatory    ContractRole = "signatory"
	RoleCounterparty ContractRole = "counterparty"
	RoleBeneficiary  ContractRole = "beneficiary"
	RoleWitness      ContractRole = "witness"
)

type ContractPerson struct {
	ContractID uuid.UUID `gorm:"type:char(36);primaryKey"`
	Contract   Contract  `gorm:"constraint:OnDelete:CASCADE;foreignKey:ContractID;references:ID"`

	PersonID uuid.UUID `gorm:"type:char(36);primaryKey"`
	Person   Person    `gorm:"constraint:OnDelete:CASCADE;foreignKey:PersonID;references:ID"`

	Role ContractRole `gorm:"type:text;not null"`
}
