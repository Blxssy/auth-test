package storage

import (
	"fmt"
	"github.com/Blxssy/auth-test/internal/config"
	"github.com/Blxssy/auth-test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type Storage interface {
	CreateToken(token *models.Token) error
	GetToken(accessToken string) (*models.Token, error)

	GetUser(email string) (*models.User, error)
}

type storage struct {
	db *gorm.DB
}

func NewStorage(logger *slog.Logger, config *config.Config) Storage {
	db, err := connectDatabase(config)
	if err != nil {
		logger.Error("Failure database connection")
	}
	logger.Info("Successfully connection to database")
	logger.Info("db", slog.String("port", config.Port))

	db.AutoMigrate(&models.User{}, models.Token{})

	return &storage{
		db: db,
	}
}

func connectDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.DBName,
		config.Database.Password)
	return gorm.Open(postgres.Open(dsn))
}

func (s *storage) CreateToken(token *models.Token) error {
	return s.db.Create(token).Error
}

func (s *storage) GetToken(accessToken string) (*models.Token, error) {
	return nil, nil
}

func (s *storage) GetUser(email string) (*models.User, error) {
	return nil, nil
}
