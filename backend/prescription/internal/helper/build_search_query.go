package helper

import (
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var validParams = []string{
	"dosage",
	"medication",
	"notes",
	"started",
	"ended",
}

func BuildQueryWithSearchParam(searchQueries map[string]string, db *gorm.DB) *gorm.DB {

	if len(searchQueries) == 0 {
		return db.Where("ended IS NULL")
	} else {
		dbChain := db
		for key, value := range searchQueries {
			if slices.Contains(validParams, key) {
				dbChain = dbChain.Where(key+" = ?", value)
			}

		}
		return dbChain
	}

}
