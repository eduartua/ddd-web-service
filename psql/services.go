package psql

import (
	app "github.com/eduartua/ddd-web-service"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StoresConfig func(*Stores) error

func WithGorm(connectionInfo string) StoresConfig {
	return func(s *Stores) error {
		DB, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{})
		if err != nil {
			return err
		}
		s.DB = DB
		return nil
	}
}

func WithLogMode(mode bool) StoresConfig {
	newLogger := logger.Default
	if !mode {
		newLogger = logger.Default.LogMode(logger.Silent)
		return func(s *Stores) error {
			s.DB.Config.Logger = newLogger
			return nil
		}
	}
	return func(s *Stores) error {
		s.DB.Config.Logger = newLogger
		return nil
	}
}

func WithUser(pepper, hmacKey string) StoresConfig {
	return func(s *Stores) error {
		s.User = NewUserStore(s.DB, pepper, hmacKey)
		return nil
	}
}

func NewStores(cfgs ...StoresConfig) (*Stores, error) {
	var s Stores
	for _, cfg := range cfgs {
		if err := cfg(&s); err != nil {
			return nil, err
		}
	}
	return &s, nil
}

type Stores struct {
	User        app.UserStore
	DB          *gorm.DB
}

// Close closes de database connection
func (s *Stores) Close() error {
	psqlDB, _ := s.DB.DB()
	return psqlDB.Close()
}

// AutoMigrate will attempt to automatically migrate all tables
func (s *Stores) AutoMigrate() error {
	return s.DB.AutoMigrate(&app.User{}, &pwReset{})
}

// DestructiveReset drops all tables and rebuilds them
func (s *Stores) DestructiveReset() error {
	err := s.DB.Migrator().DropTable(&app.User{}, &pwReset{})
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}
