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

func InsertCities(db *gorm.DB) {
	// Path folder yang berisi file JSON
	folderPath := "json/kabupaten_kota"

	// Membaca daftar file di dalam folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	re := regexp.MustCompile(`kab-(\d+)\.json`)
	var datas []domain.City

	for _, file := range files {
		// Memastikan hanya membaca file dengan ekstensi .json
		if filepath.Ext(file.Name()) == ".json" {
			matches := re.FindStringSubmatch(file.Name())
			if len(matches) > 1 {
				codeProvinces := matches[1]

				// get data province by code
				provincieID, err := GetProvinces(db, codeProvinces)
				if err != nil {
					log.Fatalf("Error getting province data: %v", err)
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

				for codeCity, name := range data {
					province := domain.City{
						Code:       codeCity,
						Name:       name,
						IsActive:   true,
						ProvinceID: provincieID.ID,
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}

					datas = append(datas, province)
				}
			}
		}
	}

	if err := db.CreateInBatches(&datas, 100).Error; err != nil {
		log.Fatalf("Error inserting province data: %v", err)
	}

	fmt.Println("Data City berhasil disimpan.")
}
