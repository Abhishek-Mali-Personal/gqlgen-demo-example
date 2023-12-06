package models

import "time"

type Task struct {
	ID          int         `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
	Name        string      `gorm:"column:name;type:varchar(255);not null;unique;" json:"name"`
	Description *string     `gorm:"column:description;type:varchar(255);" json:"description"`
	StartDate   time.Time   `gorm:"column:start_date;not null;" json:"start_date"`
	EndDate     *time.Time  `gorm:"column:end_date;" json:"end_date,omitempty"`
	Employees   []*Employee `gorm:"many2many:employee_tasks;" json:"employees"`
}

func (Task) IsResultUnion() {}
