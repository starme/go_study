package drivers

//import (
//	"fmt"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"gorm.io/gorm/schema"
//)
//
//type SqliteConnector struct {
//	TablePrefix string
//	LogLevel    logger.LogLevel
//	Dsn         string
//}
//
//func (s *SqliteConnector) Connect() gorm.Dialector {
//	fmt.Println("Sqlite connected")
//	return sqlite.Open(s.Dsn)
//}
//
//func (s *SqliteConnector) Config() *gorm.Config {
//	return &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix: s.TablePrefix,
//		},
//		Logger: logger.Default.LogMode(s.LogLevel),
//	}
//}
//
//func makeSqliteConnector(c Config) Connector {
//	return &SqliteConnector{
//		TablePrefix: c["table_prefix"].(string),
//		LogLevel:    c["log_level"].(logger.LogLevel),
//		Dsn:         c["database"].(string),
//	}
//}
