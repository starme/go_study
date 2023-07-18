package database

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"star/pkg/config"
	"star/pkg/database/drivers"
	"sync"
	"time"
)

var (
	connections Connections
	conf        *DBConfig
	dbOnce      sync.Once
)

type DBConfig struct {
	Default string
	Drivers []drivers.Config
}

func init() {
	configs := config.Sub("database")
	if configs == nil {
		panic("database config not found")
	}

	conf = &DBConfig{
		Default: configs.GetString("default"),
		Drivers: parseConfigToDriver(configs),
	}

	connections = make(Connections)
	for _, d := range conf.Drivers {
		connections.Add(d.GetDriverName(), dbFactory(d))
	}
}

func parseConfigToDriver(configs *viper.Viper) []drivers.Config {
	var c []drivers.Config
	dbDrivers, ok := configs.Get("drivers").(map[string]interface{})
	if !ok {
		panic("database config drivers not found")
	}

	for n, driver := range dbDrivers {
		d, ok := driver.(map[string]interface{})
		if !ok {
			panic("database config driver not found")
		}
		d["name"] = n
		c = append(c, d)
	}
	return c
}

func dbFactory(d drivers.Config) *gorm.DB {
	dbresolver.Register(dbresolver.Config{})
	driver := drivers.Factory(d)
	db, err := gorm.Open(driver.Connect(), driver.Config())
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(d.GetMaxIdleConns())

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(d.GetMaxOpenConns())

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func Connection(name string) *gorm.DB {
	return connections.Get(name)
}

func DB() *gorm.DB {
	return Connection(conf.Default)
}
