package db

import (
	"fmt"
	"log"
	l "log"

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
	dsn := fmt.Sprintf("host=localhost user=postgres password=example dbname=uber port=5432 application_name=dashboardapi sslmode=disable TimeZone=utc")

	gormDB, err := gorm.Open(pd.Open(dsn), &gorm.Config{
		Logger: logger.New(
			l.New(os.Stdout, "\r\n", l.LstdFlags), // io writer
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
