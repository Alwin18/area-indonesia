package domain

import "time"

type Province struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NamaProvinsi string    `gorm:"not null" json:"nama_provinsi"`
	KodeProvinsi string    `gorm:"not null, unique" json:"kode_provinsi"`
	Status       string    `gorm:"not null" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Province) TableName() string {
	return "provinsi"
}
