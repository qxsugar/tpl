package repo

import (
	"api-pkg/internal/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var dbOne sync.Once

func mustInit() {
	dbOne.Do(func() {
		cfg := config.GetConfig()
		var err error
		db, err = gorm.Open(mysql.Open(cfg.Database.Default))
		if err != nil {
			panic(fmt.Sprintf("数据库初始化失败！db_source: %s, err: %s", cfg.Database.Default, err))
		}
	})
}

func GetDatabase() *gorm.DB {
	mustInit()
	return db
}
