package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	Marca string `json:"marca"`
	Modelo string `json:"modelo"`
	Ano int `json:"ano"`
	Placa string `json:"placa"`
}
