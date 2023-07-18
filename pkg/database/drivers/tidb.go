package drivers

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"gorm.io/gorm/schema"
//)
//
//type TidbConnector struct {
//	TablePrefix string
//	LogLevel    logger.LogLevel
//	Dsn         string
//}
//
//func (t *TidbConnector) Connect() gorm.Dialector {
//	fmt.Println("Sqlite connected")
//	return mysql.Open(t.Dsn)
//}
//
//func (t *TidbConnector) Config() *gorm.Config {
//	return &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix: t.TablePrefix,
//		},
//		Logger: logger.Default.LogMode(t.LogLevel),
//	}
//}
//
//func makeTidbConnector(c Config) Connector {
//	dsn := fmt.Sprintf(
//		"%s:@tcp(%s:%d)/%s",
//		c["username"], c["host"], c["port"], c["database"],
//	)
//	return &SqlServerConnector{
//		TablePrefix: c["table_prefix"].(string),
//		LogLevel:    c["log_level"].(logger.LogLevel),
//		Dsn:         dsn,
//	}
//}
