package helper

import (
	"gorm.io/gorm"
)

func BuildQueryWithSearchParam(searchQueries map[string]string, db *gorm.DB) *gorm.DB {

	if len(searchQueries) == 0 {
		return db.Where("ended IS NULL")
	} else {
		dbChain := db
		for key, value := range searchQueries {
			dbChain = dbChain.Where(key+" = ?", value)
		}
		return dbChain
	}

}
