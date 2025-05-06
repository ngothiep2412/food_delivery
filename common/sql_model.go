package common

import "time"

type SqlModel struct {
	Id       int        `json:"id" gorm:"column:id"`
	Status   int        `json:"status" gorm:"column:status;default:1"`
	CreateAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdateAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
