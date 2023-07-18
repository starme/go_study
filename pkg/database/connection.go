package database

import (
	"gorm.io/gorm"
)

type Connections map[string]*gorm.DB

func (c Connections) Add(name string, db *gorm.DB) {
	c[name] = db
}

func (c Connections) Get(name string) *gorm.DB {
	return c[name]
}

func (c Connections) Close() {
	for name := range c {
		c.CloseByName(name)
		delete(c, name)
	}
}

func (c Connections) CloseByName(name string) {
	sqlDB, _ := c[name].DB()
	if err := sqlDB.Close(); err != nil {
		return
	}
}
