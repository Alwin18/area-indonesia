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

func InsertDistrict(db *gorm.DB) {
	// Path folder yang berisi file JSON
	folderPath := "json/kecamatan"

	// Membaca daftar file di dalam folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	re := regexp.MustCompile(`kec-(\d+)\-(\d+)\.json`)
	var datas []domain.District
	for _, file := range files {
		matches := re.FindStringSubmatch(file.Name())
		if len(matches) > 2 {
			codeProvinces := matches[1]
			codeCities := matches[2]

			// get city id by code
			cityID, err := domain.GetCities(db, codeProvinces, codeCities)
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

			for codeDistrict, name := range data {
				province := domain.District{
					KodeKecamatan: codeDistrict,
					NamaKecamatan: name,
					Status:        "Aktif",
					KotaID:        cityID,
					CreatedAt:     time.Now(),
					UpdatedAt:     time.Now(),
				}

				datas = append(datas, province)
			}
		}
	}

	if err := db.CreateInBatches(&datas, 100).Error; err != nil {
		log.Fatalf("Error inserting City data: %v", err)
	}

	fmt.Println("Data City berhasil disimpan.")
}
