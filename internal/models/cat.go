package models

import (
	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model
	Happiness int `json:"happiness"`
	Fullness  int `json:"fullness"`
	Health    int `json:"health"`
}
