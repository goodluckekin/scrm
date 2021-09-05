package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"scrm/app/auth/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewAuthRepo)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "auth-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "auth-service/data"))

	d := &Data{
		db:  db,
		log: log,
	}
	return d, func() {

	}, nil
}
