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
	ID         uuid.UUID `gorm:"type:char(36); primaryKey" json:"id"`
	ContractID uuid.UUID
	Contract   Contract `gorm:"foreignKey:ContractID"`
	PersonID   uuid.UUID
	Person     Person       `gorm:"foreignKey:PersonID"`
	Role       ContractRole `gorm:"type:text; not null" json:"role"`
}
