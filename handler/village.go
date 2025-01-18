package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"

	domain "github.com/Alwin18/indonesia-area/model"
	"gorm.io/gorm"
)

func InsertVillage(db *gorm.DB) {
	// Path folder yang berisi file JSON
	folderPath := "json/kelurahan_desa"

	// Membaca daftar file di dalam folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	re := regexp.MustCompile(`keldesa-(\d+)\-(\d+)\-(\d+)\.json`)
	var datas []domain.Village
	for _, file := range files {
		matches := re.FindStringSubmatch(file.Name())
		if len(matches) > 3 {
			codeProvinces := matches[1]
			codeCities := matches[2]
			codeDistrict := matches[3]

			// get city id by code
			districtID, err := domain.GetDistrict(db, codeProvinces, codeCities, codeDistrict)
			if err != nil {
				log.Fatalf("Error getting city data: %v", err)
			}

			filePath := filepath.Join(folderPath, file.Name())

			// Membaca file JSON
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Error reading file %s: %v", file.Name(), err)
				return
			}

			var data map[string]string
			if err := json.Unmarshal(fileData, &data); err != nil {
				log.Printf("Error unmarshaling file %s: %v", file.Name(), err)
				return
			}

			for codeVillage, name := range data {
				province := domain.Village{
					Code:       codeVillage,
					Name:       name,
					IsActive:   true,
					DistrictID: districtID,
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}

				datas = append(datas, province)
			}
		}
	}

	if err := db.CreateInBatches(&datas, 100).Error; err != nil {
		log.Fatalf("Error inserting City data: %v", err)
	}

	fmt.Println("Data Village berhasil disimpan.")
}
