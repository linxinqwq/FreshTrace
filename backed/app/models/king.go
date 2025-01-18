package models

import "gorm.io/gorm"

// FruitKind 代表水果种类
type FruitKind struct {
	gorm.Model
	Kind string `json:"kind"`
}
