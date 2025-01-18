package domain

import "time"

type Village struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	Code       string    `gorm:"not null, unique" json:"code"`
	DistrictID int64     `gorm:"not null" json:"city_id"`
	Name       string    `gorm:"not null" json:"name"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	District   District  `gorm:"foreignKey:DistrictID" json:"district"`
}

func (Village) TableName() string {
	return "villages"
}
