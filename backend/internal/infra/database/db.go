package database

import (
	"fmt"
	"sync"

	"github.com/luizrgf2/real-time-chat-go/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *Database
	once       sync.Once
)

type Database struct {
	Db *gorm.DB
}

func GetDatabaseInstance() *Database {
	once.Do(func() {
		configDb := config.LoadDatabaseConfig()

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", *configDb.DbUser, *configDb.DbPass, *configDb.DbHost, *configDb.DbPort, *configDb.DbDatabase)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		dbInstance = &Database{Db: db}
	})
	return dbInstance
}
