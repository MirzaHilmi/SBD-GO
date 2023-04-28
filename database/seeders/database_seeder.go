package seeders

import "gorm.io/gorm"

func Seed(db gorm.DB, model []interface{}) {
	db.Create(&model)
}
