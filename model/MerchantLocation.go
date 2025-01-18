package domain

import "time"

type MerchantLocation struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	MerchantID int64     `gorm:"not null" json:"merchant_id"`
	ProvinceID int64     `gorm:"not null" json:"province_id"`
	CityID     int64     `gorm:"not null" json:"city_id"`
	DistrictID int64     `gorm:"not null" json:"district_id"`
	PostalCode string    `json:"postal_code"`
	Latitude   float64   `gorm:"not null" json:"latitude"`
	Longitude  float64   `gorm:"not null" json:"longitude"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Merchant   Merchant  `gorm:"foreignKey:MerchantID" json:"merchant"`
	Province   Province  `gorm:"foreignKey:ProvinceID" json:"province"`
	City       City      `gorm:"foreignKey:CityID" json:"city"`
	District   District  `gorm:"foreignKey:DistrictID" json:"district"`
}

func (MerchantLocation) TableName() string {
	return "merchant_locations"
}
