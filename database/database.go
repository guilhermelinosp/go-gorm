package database

import (
	"log"
	"os"

	"github.com/guilhermelinosp/go-gorm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DB, err := gorm.Open(mysql.Open(os.Getenv("MYSQL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados:", err)
	}

	err = DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Falha ao migrar o modelo Product:", err)
	}
}
