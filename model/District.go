package domain

import (
	"time"

	"gorm.io/gorm"
)

type District struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	KotaID        uint      `gorm:"not null" json:"kota_id"`
	KodeKecamatan string    `gorm:"not null, unique" json:"kode_kecamatan"`
	NamaKecamatan string    `gorm:"not null" json:"nama_kecamatan"`
	Status        string    `gorm:"not null" json:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	City City `gorm:"foreignKey:KotaID" json:"city"`
}

func (District) TableName() string {
	return "kecamatan"
}

func GetDistrict(db *gorm.DB, codeProv, codeCity, codeDistrict string) (int64, error) {
	var cityID int64
	if err := db.Table("provinsi p").
		Select("d.id").
		Joins("join kota c on p.id = c.provinsi_id").
		Joins("join kecamatan d on c.id = d.kota_id").
		Where("p.kode_provinsi = ?", codeProv).
		Where("c.kode_kota = ?", codeCity).
		Where("d.kode_kecamatan = ?", codeDistrict).
		Scan(&cityID).Error; err != nil {
		return 0, err
	}
	return cityID, nil
}
