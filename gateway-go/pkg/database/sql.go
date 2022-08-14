package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type sqlDatabase struct {
	db  *gorm.DB
	dsn string
}

func NewSQLDatabase(dsn string) *sqlDatabase {

	instance := &sqlDatabase{
		dsn: dsn,
	}
	return instance
}

func (s *sqlDatabase) Connect() {
	log.Println("Connecting to database through dsn: " + s.dsn)
	db, err := gorm.Open(postgres.Open(s.dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Cannot connect to database")
	}
	s.db = db
	log.Println("Connect success to database ")
}

func (s *sqlDatabase) GetDatabase() *gorm.DB {
	return s.db
}
