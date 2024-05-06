package postgres

import (
	"github.com/zhikariz/go-commerce/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(config *configs.PostgresConfig) (*gorm.DB, error) {
	dsn := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.Database + " port=" + config.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
