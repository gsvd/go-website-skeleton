package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func New(dsn string) (*Store, error) {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(); err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

func (s *Store) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
