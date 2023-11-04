package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return
	}
	envPath := fmt.Sprintf("%s/%s", filepath.Dir(path), ".env.local")

	envFile, err := godotenv.Read(envPath)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	user := envFile["DB_USER"]
	pass := envFile["DB_PASS"]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, "localhost", "3306", "mysql")

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
