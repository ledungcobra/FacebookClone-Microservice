package database

import (
	"log"

	"ledungcobra/gateway-go/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLDBManager struct {
	db  *gorm.DB
	dsn string
}

func NewSQLDatabase(dsn string) *SQLDBManager {

	instance := &SQLDBManager{
		dsn: dsn,
	}
	return instance
}

func (s *SQLDBManager) Connect() {
	log.Println("Connecting to database through dsn: " + s.dsn)
	db, err := gorm.Open(postgres.Open(s.dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot connect to database")
	}
	s.db = db
	log.Println("Connect success to database ")
}

func (s *SQLDBManager) GetDatabase() *gorm.DB {
	return s.db
}

func (s *SQLDBManager) MigrateModels() {
	log.Println("Migrating to database")
	if err := s.db.Migrator().AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Panic("Cannot migrate models because of " + err.Error())
	}
}
