package models

import (
	"time"

	"gorm.io/gorm"
)

type Solicitor struct {
	gorm.Model

	Dni       string    `gorm:"not null" json:"dni"`
	Name      string    `gorm:"type:varchar(15);not null" json:"name"`
	LastName  string    `gorm:"type:varchar(15);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(100);not null" json:"email"`
	Role      string    `gorm:"type:varchar(25);not null" json:"role"`
	CreatedAt time.Time `gorm:"type:datetime" json:"CreatedAt"`
}
