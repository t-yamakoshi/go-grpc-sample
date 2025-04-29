package mysql

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID          int64  `gorm:"column:id;primaryKey" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	IsCompleted bool   `gorm:"column:is_completed" json:"is_completed"`
	CreatedAt   string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updated_at"`
}
