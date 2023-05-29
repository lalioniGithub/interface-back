package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model

	ID       int    `gorm:"not null;unique_index"`
	Tipo     string `gorm:"type:varchar(5);not null"`
	Flujo    string `gorm:"type:varchar(50);not null"`
	Camino   string `gorm:"type:varchar(15);not null"`
	Variante string `gorm:"not null"`
}
