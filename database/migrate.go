package database

import (
	"auth_api/entity"
	"log"

	"gorm.io/gorm"
)

func Migrate(instance *gorm.DB) {
	if err := instance.AutoMigrate(&entity.User{}); err != nil {
		log.Fatalln("Could not migrate models to database!")
	}
	log.Println("Database migration completed!")
}
