package models

type PersonType string

const (
	PersonTypeIndividual PersonType = "individual"
	PersonTypeEntity     PersonType = "entity"
)

type Person struct {
	BaseModel
	Type  PersonType `gorm:"type:text; default:'individual'; not null; index" json:"type"`
	Name  string     `gorm:"size:255" json:"name"`
	Code  string     `gorm:"size:50; uniqueIndex" json:"code"`
	Email *string    `gorm:"size:100; unique" json:"email,omitempty"`
	Phone *string    `gorm:"size:20; unique" json:"phone,omitempty"`
}
