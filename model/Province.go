package domain

import "time"

type Province struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Code      string    `gorm:"not null, unique" json:"code"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Province) TableName() string {
	return "provinces"
}
