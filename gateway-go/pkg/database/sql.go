package database

import (
	"log"

	"ledungcobra/gateway-go/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLConnector struct {
	db  *gorm.DB
	dsn string
}

func NewSQLConnector(dsn string) *SQLConnector {

	instance := &SQLConnector{
		dsn: dsn,
	}
	return instance
}

func (s *SQLConnector) Connect() {
	log.Println("Connecting to database through dsn: " + s.dsn)
	db, err := gorm.Open(postgres.Open(s.dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot connect to database")
	}
	s.db = db
	log.Println("Connect success to database ")
}

func (s *SQLConnector) GetDatabase() *gorm.DB {
	return s.db
}

func (s *SQLConnector) MigrateModels() {
	log.Println("Migrating to database")
	if err := s.db.Migrator().AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		log.Panic("Cannot migrate models because of " + err.Error())
	}
}
