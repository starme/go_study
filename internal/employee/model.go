package employee

import "star/pkg/database"

type Employee struct {
	database.Model

	WorkCode string `gorm:"column:workcode;type:varchar(255);not null;unique"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`

	database.Timestamps
	//database.SoftDelete
}

// TableName gives table name of model
func (e Employee) TableName() string {
	return "admin_employees"
}
