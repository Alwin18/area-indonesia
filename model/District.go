package domain

import (
	"time"

	"gorm.io/gorm"
)

type District struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	CityID    int64     `gorm:"not null" json:"city_id"`
	Code      string    `gorm:"not null, unique" json:"code"`
	Name      string    `gorm:"not null" json:"name"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	City      City      `gorm:"foreignKey:CityID" json:"city"`
}

func (District) TableName() string {
	return "districts"
}

func GetDistrict(db *gorm.DB, codeProv, codeCity, codeDistrict string) (int64, error) {
	var cityID int64
	if err := db.Table("provinces p").
		Select("d.id").
		Joins("join cities c on p.id = c.province_id").
		Joins("join districts d on c.id = d.city_id").
		Where("p.code = ?", codeProv).
		Where("c.code = ?", codeCity).
		Where("d.code = ?", codeDistrict).
		Scan(&cityID).Error; err != nil {
		return 0, err
	}
	return cityID, nil
}
