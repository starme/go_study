package drivers

//import (
//	"fmt"
//	"gorm.io/driver/sqlserver"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"gorm.io/gorm/schema"
//)
//
//type SqlServerConnector struct {
//	TablePrefix string
//	LogLevel    logger.LogLevel
//	Dsn         string
//}
//
//func (s *SqlServerConnector) Connect() gorm.Dialector {
//	fmt.Println("Sqlite connected")
//	return sqlserver.Open(s.Dsn)
//}
//
//func (s *SqlServerConnector) Config() *gorm.Config {
//	return &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix: s.TablePrefix,
//		},
//		Logger: logger.Default.LogMode(s.LogLevel),
//	}
//}
//
//func makeSqlServerConnector(c Config) Connector {
//	dsn := fmt.Sprintf(
//		"sqlserver://%s:%s@%s:%s?database=%s",
//		c["username"], c["password"], c["host"], c["port"], c["database"],
//	)
//	return &SqlServerConnector{
//		TablePrefix: c["table_prefix"].(string),
//		LogLevel:    c["log_level"].(logger.LogLevel),
//		Dsn:         dsn,
//	}
//}
//
