package domain

import "time"

type MerchantFloor struct {
	ID             int64           `gorm:"primaryKey" json:"id"`
	MerchantID     int64           `gorm:"not null" json:"merchant_id"`
	FloorNumber    int             `gorm:"not null;unique" json:"floor_number"`
	FloorName      string          `gorm:"not null" json:"floor_name"`
	IsRooftop      bool            `gorm:"default:false" json:"is_rooftop"`
	Description    string          `json:"description"`
	IsFull         bool            `gorm:"default:false" json:"is_full"`
	IsActive       bool            `gorm:"default:true" json:"is_active"`
	PictureURL     string          `gorm:"not null" json:"picture_url"`
	CreatedAt      time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	Merchant       Merchant        `gorm:"foreignKey:MerchantID" json:"merchant"`
	MerchantTables []MerchantTable `gorm:"foreignKey:MerchantFloorID" json:"merchant_tables"`
}

func (MerchantFloor) TableName() string {
	return "merchant_floors"
}
