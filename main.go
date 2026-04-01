package main

import (
	"fmt"
	"runtime"

	"github.com/Alwin18/indonesia-area/database"
	"github.com/Alwin18/indonesia-area/handler"
)

func main() {
	fmt.Printf("Available CPU cores: %d\n", runtime.NumCPU())

	runtime.GOMAXPROCS(5)

	db := database.NewDatabase()

	// insert province
	// handler.InsertProvince(db)

	// insert city
	// handler.InsertCities(db)

	// insert District
	handler.InsertDistrict(db)

	// insert Village
	// handler.InsertVillage(db)
}
