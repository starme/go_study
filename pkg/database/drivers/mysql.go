package drivers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MysqlConnector struct {
	TablePrefix string
	LogLevel    logger.LogLevel
	Dsn         string
}

func (m *MysqlConnector) Connect() gorm.Dialector {
	fmt.Println("Mysql connected")
	return mysql.Open(m.Dsn)
}

func (m *MysqlConnector) Config() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: m.TablePrefix,
		},
		Logger: logger.Default.LogMode(m.LogLevel),
	}
}

func makeMysqlConnector(c Config) Connector {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c["username"], c["password"], c["host"], c["port"], c["database"], c["charset"],
	)

	return &MysqlConnector{
		TablePrefix: c["table_prefix"].(string),
		LogLevel:    logger.LogLevel(c["log_level"].(int)),
		Dsn:         dsn,
	}
}
