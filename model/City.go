package domain

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	ProvinceID int64     `gorm:"not null" json:"province_id"`
	Code       string    `gorm:"not null, unique" json:"code"`
	Name       string    `gorm:"not null" json:"name"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (City) TableName() string {
	return "cities"
}

func GetCities(db *gorm.DB, codeProv, codeCity string) (int64, error) {
	var cityID int64
	if err := db.Table("provinces p").
		Select("c.id").
		Joins("join cities c on p.id = c.province_id").
		Where("p.code = ?", codeProv).
		Where("c.code = ?", codeCity).
		Scan(&cityID).Error; err != nil {
		return 0, err
	}
	return cityID, nil
}
