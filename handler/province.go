package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	domain "github.com/Alwin18/indonesia-area/model"
	"gorm.io/gorm"
)

func InsertProvince(db *gorm.DB) {
	// Membuka file JSON
	file, err := os.Open("json/provinsi/provinsi.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Mengambil data JSON
	var provinces map[string]string
	if err := json.NewDecoder(file).Decode(&provinces); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Menyimpan data provinsi ke database
	var datas []domain.Province
	for code, name := range provinces {
		province := domain.Province{
			Code:      code,
			Name:      name,
			IsActive:  true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		datas = append(datas, province)
	}

	if err := db.Create(&datas).Error; err != nil {
		log.Fatalf("Error inserting province data: %v", err)
	}

	fmt.Println("Data provinces berhasil disimpan.")
}

func GetProvinces(db *gorm.DB, code string) (*domain.Province, error) {
	var provinces domain.Province

	if err := db.Select("id").Where("code = ?", code).First(&provinces).Error; err != nil {
		return nil, err
	}

	return &provinces, nil
}
