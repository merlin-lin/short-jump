package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"short-jump/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDBEngine, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

func NewDBEngine(c *conf.Data) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
}
