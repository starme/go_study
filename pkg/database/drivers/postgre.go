package drivers

//import (
//	"fmt"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"gorm.io/gorm/schema"
//)
//
//type PostgreConnector struct {
//	TablePrefix          string
//	LogLevel             logger.LogLevel
//	Dsn                  string
//	PreferSimpleProtocol bool
//}
//
//func (p *PostgreConnector) Connect() gorm.Dialector {
//	fmt.Println("Sqlite connected")
//	if p.PreferSimpleProtocol {
//		return postgres.New(postgres.Config{
//			DSN:                  p.Dsn,
//			PreferSimpleProtocol: p.PreferSimpleProtocol,
//		})
//	}
//	return postgres.Open(p.Dsn)
//}
//
//func (p *PostgreConnector) Config() *gorm.Config {
//	return &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix: p.TablePrefix,
//		},
//		Logger: logger.Default.LogMode(p.LogLevel),
//	}
//}
//
//func makePostgreConnector(c Config) Connector {
//	dsn := fmt.Sprintf(
//		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
//		c["host"], c["username"], c["password"], c["database"], c["port"], c["timezone"],
//	)
//	return &PostgreConnector{
//		TablePrefix:          c["table_prefix"].(string),
//		LogLevel:             c["log_level"].(logger.LogLevel),
//		Dsn:                  dsn,
//		PreferSimpleProtocol: c["prefer_simple_protocol"].(bool),
//	}
//}
