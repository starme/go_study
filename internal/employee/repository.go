package employee

import "star/pkg/database"

func List() []Employee {
	var employees []Employee
	database.DB().Limit(20).Find(&employees)
	return employees
}
