package helper

import (
	"strconv"

	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var validParams = []string{
	"dosage",
	"medication",
	"notes",
	"started",
	"ended",
	"present",
}

func BuildQueryWithSearchParam(searchQueries map[string]string, db *gorm.DB) *gorm.DB {

	dbChain := db

	for key, value := range searchQueries {
		if slices.Contains(validParams, key) {
			if key == "present" {
				isPresent, _ := strconv.ParseBool(value)
				if isPresent {
					dbChain = dbChain.Where("ended is NULL")
				} else {
					dbChain = dbChain.Where("ended is NOT NULL")
				}
			} else {
				dbChain = dbChain.Where(key+" = ?", value)
			}
		}
	}

	return dbChain

}
