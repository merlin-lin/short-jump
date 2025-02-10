package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"short-jump/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDbEngine)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

func NewDbEngine(c *conf.Data) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	return db, err
}
