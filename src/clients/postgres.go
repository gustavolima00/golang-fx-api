package clients

import (
	"fmt"
	"go-api/src/common/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go.uber.org/fx"
)

type PostgresClient interface {
	GetConnection() *gorm.DB
}

type PostgresClientParams struct {
	fx.In

	Config *config.Config
}

type postgresClient struct {
	db     *gorm.DB
	config *config.Config
}

func NewPostgresClient(p PostgresClientParams) (PostgresClient, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		p.Config.PostgresHost,
		p.Config.PostgresUser,
		p.Config.PostgresPassword,
		p.Config.PostgresDB,
		p.Config.PostgresPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &postgresClient{
		db:     db,
		config: p.Config,
	}, nil
}

func (c *postgresClient) GetConnection() *gorm.DB {
	return c.db
}
