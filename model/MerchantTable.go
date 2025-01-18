package domain

import "time"

type MerchantTable struct {
	ID              int64         `gorm:"primaryKey" json:"id"`
	MerchantID      int64         `gorm:"not null" json:"merchant_id"`
	MerchantFloorID int64         `gorm:"not null" json:"merchant_floor_id"`
	TableNumber     int           `gorm:"unique" json:"table_number"`
	IsReserved      bool          `gorm:"default:false" json:"is_reserved"`
	IsActive        bool          `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	Merchant        Merchant      `gorm:"foreignKey:MerchantID" json:"merchant"`
	MerchantFloor   MerchantFloor `gorm:"foreignKey:MerchantFloorID" json:"merchant_floor"`
}

func (MerchantTable) TableName() string {
	return "merchant_tables"
}
