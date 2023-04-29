package seeders

import (
	"github.com/MirzaHilmi/SBD-GO/models"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, model *[]models.Address) error {
	if err := db.Create(&model); err.Error != nil {
		return err.Error
	}

	return nil
}
