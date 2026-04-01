package domain

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ProvinsiID uint      `gorm:"not null" json:"provinsi_id"`
	KodeKota   string    `gorm:"not null, unique" json:"kode_kota"`
	NamaKota   string    `gorm:"not null" json:"nama_kota"`
	Status     string    `gorm:"not null" json:"status"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (City) TableName() string {
	return "kota"
}

func GetCities(db *gorm.DB, codeProv, codeCity string) (uint, error) {
	var cityID uint
	if err := db.Table("provinsi p").
		Select("c.id").
		Joins("join kota c on p.id = c.provinsi_id").
		Where("p.kode_provinsi = ?", codeProv).
		Where("c.kode_kota = ?", codeCity).
		Scan(&cityID).Error; err != nil {
		return 0, err
	}
	return cityID, nil
}
