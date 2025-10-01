package migration

import (
	"log"
	"ms-golang-fiber/internal/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("Migration success")
}
