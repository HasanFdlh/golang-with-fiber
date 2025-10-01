package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitPostgres() *gorm.DB {
	host := viper.GetString("POSTGRE_HOST")
	port := viper.GetString("POSTGRE_PORT")
	user := viper.GetString("POSTGRE_USER")
	pass := viper.GetString("POSTGRE_PASSWORD")
	name := viper.GetString("POSTGRE_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect POSTGRE: %v", err)
	}
	DB = db
	return db
}

func InitMysql() *gorm.DB {
	host := viper.GetString("MYSQL_HOST")
	port := viper.GetString("MYSQL_PORT")
	user := viper.GetString("MYSQL_USER")
	password := viper.GetString("MYSQL_PASSWORD")
	name := viper.GetString("MYSQL_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect MySQL: %v", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ MySQL connected!")
	DB = db
	return db
}
