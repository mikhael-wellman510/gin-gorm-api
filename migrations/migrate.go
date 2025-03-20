package migrations

import (
	"GinGonicGorm/entity"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {

	err := db.AutoMigrate(&entity.Product{}, &entity.User{})

	if err != nil {

		return err
	}
	log.Println("Succes Migration !!")
	return nil
}
