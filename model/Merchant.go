package domain

import "time"

type Merchant struct {
	ID           int64              `gorm:"primaryKey" json:"id"`
	Name         string             `gorm:"not null" json:"name"`
	Address      string             `gorm:"not null" json:"address"`
	ThumbnailURL string             `gorm:"not null" json:"thumbnail_url"`
	PictureURL   string             `gorm:"not null" json:"picture_url"`
	IsActive     bool               `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
	Locations    []MerchantLocation `gorm:"foreignKey:MerchantID" json:"locations"`
	Floors       []MerchantFloor    `gorm:"foreignKey:MerchantID" json:"floors"`
}

func (Merchant) TableName() string {
	return "merchants"
}
