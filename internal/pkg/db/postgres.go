package db

import (
	"fmt"
	"log"

	"github.com/akselarzuman/go-jaeger/internal/pkg/persistence/models"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	pd "gorm.io/driver/postgres"

	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresClient struct {
	GormDB *gorm.DB
}

func NewPostgresClient() *PostgresClient {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=uber port=5432 application_name=uberapi sslmode=disable TimeZone=utc", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	gormDB, err := gorm.Open(pd.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      false,       // Disable color
			},
		),
	})

	if err != nil {
		return nil
	}

	if err := gormDB.Use(otelgorm.NewPlugin()); err != nil {
		log.Println(err.Error())
	}

	return &PostgresClient{
		GormDB: gormDB,
	}
}

func (m *PostgresClient) AutoMigrate() error {
	return m.GormDB.AutoMigrate(&models.UserPostgresModel{})
}
