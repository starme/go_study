package drivers

import "gorm.io/gorm"

type Connector interface {
	Connect() gorm.Dialector
	Config() *gorm.Config
}

type Config map[string]interface{}

func (c Config) GetDriverName() string {
	return c["name"].(string)
}

func (c Config) GetMaxIdleConns() int {
	if n, ok := c["max_idle_conn_num"].(int); ok {
		return n
	}
	return 10
}

func (c Config) GetMaxOpenConns() int {
	if n, ok := c["max_open_conn_num"].(int); ok {
		return n
	}
	return 100
}
