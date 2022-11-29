package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id"`
	Status    int        `json:"status" gorm:"status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}
